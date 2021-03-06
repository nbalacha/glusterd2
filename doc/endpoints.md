
<!---
This file is generated using commands described below. DO NOT EDIT.

$ curl -o endpoints.json -s -X GET http://127.0.0.1:24007/endpoints
$ go build pkg/tools/generate-doc.go
$ ./generate-doc
-->

# REST API Endpoints Reference

**Note:** Fields in request structs marked with "omitempty" struct tag are optional.

Name | Methods | Path | Request | Response
--- | --- | --- | --- | ---
GetVersion | GET | /version | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [VersionResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VersionResp)
VolumeCreate | POST | /volumes | [VolCreateReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolCreateReq) | [VolumeCreateResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeCreateResp)
VolumeExpand | POST | /volumes/{volname}/expand | [VolExpandReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolExpandReq) | [VolumeExpandResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeExpandResp)
VolumeOptions | POST | /volumes/{volname}/options | [VolOptionReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolOptionReq) | [VolumeOptionResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeOptionResp)
VolumeReset | DELETE | /volumes/{volname}/options | [VolOptionResetReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolOptionResetReq) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
OptionGroupList | GET | /volumes/options-group | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [OptionGroupListResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#OptionGroupListResp)
OptionGroupCreate | POST | /volumes/options-group | [OptionGroupReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#OptionGroupReq) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
OptionGroupDelete | DELETE | /volumes/options-group/{groupname} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
VolumeDelete | DELETE | /volumes/{volname} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
VolumeInfo | GET | /volumes/{volname} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [VolumeGetResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeGetResp)
VolumeBricksStatus | GET | /volumes/{volname}/bricks | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [BricksStatusResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#BricksStatusResp)
VolumeStatus | GET | /volumes/{volname}/status | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [VolumeStatusResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeStatusResp)
VolumeList | GET | /volumes | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [VolumeListResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeListResp)
VolumeStart | POST | /volumes/{volname}/start | [VolumeStartReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeStartReq) | [VolumeStartResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeStartResp)
VolumeStop | POST | /volumes/{volname}/stop | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [VolumeStopResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeStopResp)
Statedump | POST | /volumes/{volname}/statedump | [VolStatedumpReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolStatedumpReq) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
VolfilesGenerate | POST | /volfiles | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
VolfilesGet | GET | /volfiles | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
VolfilesGet | GET | /volfiles/{volfileid:.*} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
EditVolume | POST | /volumes/{volname}/edit | [VolEditReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolEditReq) | [VolumeEditResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#VolumeEditResp)
GetPeer | GET | /peers/{peerid} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [PeerGetResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#PeerGetResp)
GetPeers | GET | /peers | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [PeerListResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#PeerListResp)
DeletePeer | DELETE | /peers/{peerid} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
AddPeer | POST | /peers | [PeerAddReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#PeerAddReq) | [PeerAddResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#PeerAddResp)
EditPeer | POST | /peers/{peerid} | [PeerEditReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#PeerEditReq) | [PeerEditResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#PeerEditResp)
SetGlobalOptions | POST | /cluster/options | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
GetGlobalOptions | GET | /cluster/options/ | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
GeoreplicationCreate | POST | /geo-replication/{mastervolid}/{remotevolid} | [GeorepCreateReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepCreateReq) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepSession)
GeoreplicationStart | POST | /geo-replication/{mastervolid}/{remotevolid}/start | [GeorepCommandsReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepCommandsReq) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepSession)
GeoreplicationStop | POST | /geo-replication/{mastervolid}/{remotevolid}/stop | [GeorepCommandsReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepCommandsReq) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepSession)
GeoreplicationDelete | DELETE | /geo-replication/{mastervolid}/{remotevolid} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
GeoreplicationPause | POST | /geo-replication/{mastervolid}/{remotevolid}/pause | [GeorepCommandsReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepCommandsReq) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepSession)
GeoreplicationResume | POST | /geo-replication/{mastervolid}/{remotevolid}/resume | [GeorepCommandsReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepCommandsReq) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepSession)
GeoreplicationStatus | GET | /geo-replication/{mastervolid}/{remotevolid} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepSession)
GeoReplicationConfigGet | GET | /geo-replication/{mastervolid}/{remotevolid}/config | [GeorepOption](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepOption) | [GeorepOption](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepOption)
GeoReplicationConfigSet | POST | /geo-replication/{mastervolid}/{remotevolid}/config | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
GeoReplicationConfigReset | DELETE | /geo-replication/{mastervolid}/{remotevolid}/config | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
GeoreplicationStatusList | GET | /geo-replication | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [GeorepSession](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepSession)
GeoreplicationSshKeyGenerate | POST | /ssh-key/{volname}/generate | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [GeorepSSHPublicKey](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepSSHPublicKey)
GeoreplicationSshKeyPush | POST | /ssh-key/{volname}/push | [GeorepSSHPublicKey](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepSSHPublicKey) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
GeoreplicationSshKeyGet | GET | /ssh-key/{volname} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [GeorepSSHPublicKey](https://godoc.org/github.com/gluster/glusterd2/pkg/api#GeorepSSHPublicKey)
BitrotEnable | POST | /volumes/{volname}/bitrot/enable | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
BitrotDisable | POST | /volumes/{volname}/bitrot/disable | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
ScrubOndemand | POST | /volumes/{volname}/bitrot/scrubondemand | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
ScrubStatus | GET | /volumes/{volname}/bitrot/scrubstatus | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
QuotaEnable | POST | /quota/{volname} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
QuotaDisable | DELETE | /quota/{volname} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
QuotaList | GET | /quota/{volname}/limit | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
QuotaLimit | POST | /quota/{volname}/limit | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
QuotaRemove | DELETE | /quota/{volname}/limit | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
WebhookAdd | POST | /events/webhook | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
WebhookDelete | DELETE | /events/webhook | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
WebhookList | GET | /events/webhook | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
EventsList | GET | /events | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
GlustershEnable | POST | /volumes/{name}/heal/enable | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
GlustershDisable | POST | /volumes/{name}/heal/disable | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
DeviceAdd | POST | /devices/{peerid} | [AddDeviceReq](https://godoc.org/github.com/gluster/glusterd2/pkg/api#AddDeviceReq) | [AddDeviceResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#AddDeviceResp)
DeviceList | GET | /devices/{peerid} | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [ListDeviceResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#ListDeviceResp)
Statedump | GET | /statedump | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
List Endpoints | GET | /endpoints | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [ListEndpointsResp](https://godoc.org/github.com/gluster/glusterd2/pkg/api#ListEndpointsResp)
Glusterd2 service status | GET | /ping | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#) | [](https://godoc.org/github.com/gluster/glusterd2/pkg/api#)
