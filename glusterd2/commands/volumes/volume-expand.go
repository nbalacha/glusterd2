package volumecommands

import (
	"net/http"
	"path/filepath"

	"github.com/gluster/glusterd2/glusterd2/events"
	"github.com/gluster/glusterd2/glusterd2/gdctx"
	"github.com/gluster/glusterd2/glusterd2/peer"
	restutils "github.com/gluster/glusterd2/glusterd2/servers/rest/utils"
	"github.com/gluster/glusterd2/glusterd2/transaction"
	"github.com/gluster/glusterd2/glusterd2/volume"
	"github.com/gluster/glusterd2/pkg/api"
	"github.com/gluster/glusterd2/pkg/errors"

	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
)

func registerVolExpandStepFuncs() {
	var sfs = []struct {
		name string
		sf   transaction.StepFunc
	}{
		{"vol-expand.ValidateAndPrepare", expandValidatePrepare},
		{"vol-expand.ValidateBricks", validateBricks},
		{"vol-expand.InitBricks", initBricks},
		{"vol-expand.UndoInitBricks", undoInitBricks},
		{"vol-expand.StartBrick", startBricksOnExpand},
		{"vol-expand.UndoStartBrick", undoStartBricksOnExpand},
		{"vol-expand.UpdateVolinfo", updateVolinfoOnExpand},
		{"vol-expand.NotifyClients", notifyVolfileChange},
	}
	for _, sf := range sfs {
		transaction.RegisterStepFunc(sf.sf, sf.name)
	}
}
func checkDupBrickEntryVolExpand(req api.VolExpandReq) error {
	dupEntry := map[string]bool{}

	for _, brick := range req.Bricks {
		if dupEntry[brick.PeerID+filepath.Clean(brick.Path)] == true {
			return errors.ErrDuplicateBrickPath
		}
		dupEntry[brick.PeerID+filepath.Clean(brick.Path)] = true

	}

	return nil
}

func volumeExpandHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	logger := gdctx.GetReqLogger(ctx)
	volname := mux.Vars(r)["volname"]

	var req api.VolExpandReq
	if err := restutils.UnmarshalRequest(r, &req); err != nil {
		restutils.SendHTTPError(ctx, w, http.StatusBadRequest, errors.ErrJSONParsingFailed)
		return
	}

	if err := checkDupBrickEntryVolExpand(req); err != nil {
		restutils.SendHTTPError(ctx, w, http.StatusBadRequest, err)
		return
	}

	volinfo, err := volume.GetVolume(volname)
	if err != nil {
		restutils.SendHTTPError(ctx, w, http.StatusInternalServerError, err)
		return
	}

	for index := range volinfo.Subvols {
		for _, brick := range volinfo.Subvols[index].Bricks {

			for _, b := range req.Bricks {

				if brick.PeerID.String() == b.PeerID && brick.Path == filepath.Clean(b.Path) {
					restutils.SendHTTPError(ctx, w, http.StatusBadRequest, errors.ErrDuplicateBrickPath)
					return
				}
			}

		}
	}

	lock, unlock, err := transaction.CreateLockSteps(volname)
	if err != nil {
		restutils.SendHTTPError(ctx, w, http.StatusInternalServerError, err)
		return
	}

	txn := transaction.NewTxn(ctx)
	defer txn.Cleanup()

	nodes, err := req.Nodes()
	if err != nil {
		logger.WithError(err).Error("could not prepare node list")
		restutils.SendHTTPError(ctx, w, http.StatusInternalServerError, err)
		return
	}

	allNodes, err := peer.GetPeerIDs()
	if err != nil {
		restutils.SendHTTPError(ctx, w, http.StatusInternalServerError, err)
		return
	}

	txn.Nodes = allNodes
	txn.Steps = []*transaction.Step{
		// TODO: This is a lot of steps. We can combine a few if we
		// do not re-use the same step functions across multiple
		// volume operations.
		lock,
		{
			DoFunc: "vol-expand.ValidateAndPrepare",
			Nodes:  []uuid.UUID{gdctx.MyUUID},
		},
		{
			DoFunc: "vol-expand.ValidateBricks",
			Nodes:  nodes,
		},
		{
			DoFunc:   "vol-expand.InitBricks",
			UndoFunc: "vol-expand.UndoInitBricks",
			Nodes:    nodes,
		},
		{
			DoFunc: "vol-expand.UpdateVolinfo",
			Nodes:  []uuid.UUID{gdctx.MyUUID},
		},
		{
			DoFunc:   "vol-expand.StartBrick",
			Nodes:    nodes,
			UndoFunc: "vol-expand.UndoStartBrick",
		},
		{
			DoFunc: "vol-expand.NotifyClients",
			Nodes:  allNodes,
		},
		unlock,
	}

	if err := txn.Ctx.Set("req", &req); err != nil {
		restutils.SendHTTPError(ctx, w, http.StatusInternalServerError, err)
		return
	}

	if err := txn.Ctx.Set("volname", volname); err != nil {
		restutils.SendHTTPError(ctx, w, http.StatusInternalServerError, err)
		return
	}

	if err = txn.Do(); err != nil {
		logger.WithError(err).WithField("volume-name", volname).Error("volume expand transaction failed")
		if err == transaction.ErrLockTimeout {
			restutils.SendHTTPError(ctx, w, http.StatusConflict, err)
		} else {
			restutils.SendHTTPError(ctx, w, http.StatusInternalServerError, err)
		}
		return
	}

	volinfo, err = volume.GetVolume(volname)
	if err != nil {
		restutils.SendHTTPError(ctx, w, http.StatusInternalServerError, err)
		return
	}

	logger.WithField("volume-name", volinfo.Name).Info("volume expanded")
	events.Broadcast(volume.NewEvent(volume.EventVolumeExpanded, volinfo))

	resp := createVolumeExpandResp(volinfo)
	restutils.SendHTTPResponse(ctx, w, http.StatusOK, resp)
}

func createVolumeExpandResp(v *volume.Volinfo) *api.VolumeExpandResp {
	return (*api.VolumeExpandResp)(volume.CreateVolumeInfoResp(v))
}
