package compute

import (
	"context"

	"vendor.old2/github.com/docker/docker/api/types"
	"vendor.old2/github.com/docker/docker/api/types/network"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/rancher/agent/model"
)

func setupPublishPorts(hostConfig *container.HostConfig, instance model.Instance) {}

func setupDNSSearch(hostConfig *container.HostConfig, instance model.Instance) error {
	return nil
}

func setupLinks(hostConfig *container.HostConfig, instance model.Instance) {}

func setupNetworking(instance model.Instance, host model.Host, config *container.Config, hostConfig *container.HostConfig, client *client.Client) error {
	return nil
}

func setupFieldsHostConfig(fields model.InstanceFields, hostConfig *container.HostConfig) {

	hostConfig.LogConfig.Type = fields.LogConfig.Driver

	hostConfig.LogConfig.Config = fields.LogConfig.Config

	hostConfig.Isolation = fields.Isolation

	hostConfig.RestartPolicy = fields.RestartPolicy

	hostConfig.ConsoleSize = fields.ConsoleSize

	hostConfig.CPUCount = fields.CPUCount

	hostConfig.CPUPercent = fields.CPUPercent

	hostConfig.IOMaximumIOps = fields.IOMaximumIOps

	hostConfig.IOMaximumBandwidth = fields.IOMaximumBandwidth
}

func setupDeviceOptions(hostConfig *container.HostConfig, instance model.Instance, infoData model.InfoData) {
}

func setupComputeResourceFields(hostConfig *container.HostConfig, instance model.Instance) {

}

func dockerContainerCreate(ctx context.Context, dockerClient *client.Client, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (types.ContainerCreateResponse, error) {
	return dockerClient.ContainerCreate(context.Background(), config, hostConfig, networkingConfig, containerName)
}
