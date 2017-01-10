// +build linux freebsd solaris openbsd darwin

package storage

import (
	"bytes"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/docker/go-connections/sockets"
	"github.com/pkg/errors"
	"github.com/rancher/agent/model"
	"net/http"
	"path/filepath"
)

const (
	pluginSockDir = "/run/docker/plugins"
)

func CallVolumeDriverAttach(volume model.Volume) error {
	driver := volume.Data.Fields.Driver
	sockFile := filepath.Join(pluginSockDir, driver+".sock")
	transport := new(http.Transport)
	sockets.ConfigureTransport(transport, "unix", sockFile)
	client := &http.Client{
		Transport: transport,
	}
	url := "http://volume-plugin/VolumeDriver.Attach"

	bs, err := json.Marshal(struct{ Name string }{Name: volume.Name})
	if err != nil {
		return errors.Wrap(err, "Failed to marshal JSON")
	}

	resp, err := client.Post(url, "application/json", bytes.NewReader(bs))
	if err != nil {
		logrus.Warnf("Failed to call /VolumeDriver.Attach '%s' (driver '%s'): %s", volume.Name, driver, err)
		return nil
	}

	switch {
	case resp.StatusCode >= 200 && resp.StatusCode < 300:
		logrus.Infof("Success: /VolumeDriver.Attach '%s' (driver '%s')", volume.Name, driver)
	case resp.StatusCode >= 400 && resp.StatusCode < 500:
		logrus.Infof("/VolumeDriver.Attach '%s' is not supported by driver '%s'", volume.Name, driver)
	default:
		return errors.Errorf("/VolumeDriver.Attach '%s' (driver '%s') returned status %v: %s", volume.Name, driver, resp.StatusCode, resp.Status)
	}

	return nil
}
