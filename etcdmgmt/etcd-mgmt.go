package etcdmgmt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"
)

// ExecName is to indicate the executable name, useful for mocking in tests
var ExecName = "etcd"

func checkHealth(val time.Duration, listenClientUrls string) bool {
	result := struct{ Health string }{}
	// Checking health of etcd. Health of the etcd should be true,
	// means etcd have initialized properly before using any etcd command
	timer := time.NewTimer(time.Second * val)
	for {
		// Waiting for 15 second. Within 15 second health of etcd should
		// be true otherwise it should throw an error
		go func() {
			<-timer.C
			if result.Health != "true" {
				log.Fatal("Health of etcd is not proper. Check etcd configuration.")
			}
		}()

		resp, err := http.Get(listenClientUrls + "/health")
		if err != nil {
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal([]byte(body), &result)
		if err != nil {
			continue
		}
		if result.Health == "true" {
			timer.Stop()
			break
		}
	}
	return true
}

// StartEtcd () is to bring up etcd instance
func StartEtcd() (*exec.Cmd, error) {
	result := struct{ Health string }{}

	log.WithField("Executable", ExecName).Info("Starting")
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Could not able to get hostname")
		return nil, err
	}

	listenClientUrls := "http://" + hostname + ":2379"

	advClientUrls := "http://" + hostname + ":2379"

	listenPeerUrls := "http://" + hostname + ":2380"

	initialAdvPeerUrls := "http://" + hostname + ":2380"

	// Checking health of etcd cluster. If etcd daemon already running then
	// no need to start etcd daemon again.
	resp, err := http.Get(listenClientUrls + "/health")
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal([]byte(body), &result)
		if err == nil && result.Health == "true" {
			log.Info("etcd daemon is already running")
			return nil
		}
	}

	etcdCmd := exec.Command(ExecName,
		"-listen-client-urls", listenClientUrls,
		"-advertise-client-urls", advClientUrls,
		"-listen-peer-urls", listenPeerUrls,
		"-initial-advertise-peer-urls", initialAdvPeerUrls,
		"--initial-cluster", "default="+listenPeerUrls)

	// Don't kill chlid process (etcd) upon ^C (SIGINT) of main glusterd process
	etcdStart.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	err = etcdCmd.Start()
	if err != nil {
		log.WithField("error", err.Error()).Error("Could not start etcd daemon.")
		return nil, err
	}

	if check := checkHealth(15, listenClientUrls); check != true {
		fmt.Println("check is: ", check)
		log.Fatal("Health of etcd is not proper. Check etcd configuration.")
	}

	return etcdCmd, nil
}
