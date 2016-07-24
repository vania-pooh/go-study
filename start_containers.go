package main

import (
	cl "github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
	"os"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/network"
	"github.com/docker/go-connections/nat"
	"fmt"
)

func main() {
	hostName, hostnameError := os.Hostname()
	
	listenPort := 4444
	nodePort := 5555
	
	// Don't need to expose 4243 and no tcp connection - faster!
	dockerApiUrl := "unix:///var/run/docker.sock"
	
	dockerApiVersion := "v1.22"
	
	imageName := "selenium/node-chrome"
	defaultHeaders := map[string]string{}
	
	client, clientInitError := cl.NewClient(dockerApiUrl, dockerApiVersion, nil, defaultHeaders)

	//Pulling image from hub.docker.com
	imagePullResponse, imagePullError := client.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	
	containerConfig := container.Config{
		Image: imageName,
		
	}
	
	hostConfig := container.HostConfig{
		Privileged: true,
		PortBindings: nat.PortMap{
			fmt.Sprintf("%d/tcp", nodePort): nat.PortBinding{
				HostIP: "0.0.0.0",
				HostPort: nodePort,
			},
		},
		NetworkMode: "host",
	}
	
	networkConfig := network.NetworkingConfig{
		
	}
	
	containerCreateResponse, containerCreateError := client.ContainerCreate(context.Background(), containerConfig, networkConfig, hostConfig, "my-container")
	
	client.ContainerStart(context.Background(), containerCreateResponse.ID, types.ContainerStartOptions{})
	
}
