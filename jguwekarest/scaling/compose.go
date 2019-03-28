package main

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

var scale = 0
var file = os.Getenv("COMPOSE_FILE")
var mutex = sync.Mutex{}

func up() {
	mutex.Lock()
	newScale := strconv.Itoa(scale + 1)
	cmd := exec.Command("docker-compose", "-f", file, "up", "-d", "--scale" , "api=" + newScale)

	err := cmd.Run()
	if err != nil {
		print(err)
		mutex.Unlock()
		return
	}

	scale += 1

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key: "name",
			Value: "jguwekarest_api_" + newScale,
		}),
	})
	if err == nil {
		go getContainerStats(containers[0])
		time.Sleep(10 * time.Second)
	}

	mutex.Unlock()
}

func down() {
	mutex.Lock()
	if scale == 1 {
		mutex.Unlock()
		return
	}

	newScale := strconv.Itoa(scale - 1)
	cmd := exec.Command("docker-compose", "-f", file, "up", "-d", "--scale" , "api=" + newScale)

	err := cmd.Run()
	if err != nil {
		print(err.Error())
		mutex.Unlock()
		return
	}

	scale -= 1
	time.Sleep(10 * time.Second)
	mutex.Unlock()
}