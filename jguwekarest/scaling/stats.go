package main

import (
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"io"
	"time"
)

func listContainer() error {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key: "name",
			Value: "api",
		}),
	})

	if err != nil {
		return err
	}

	scale = len(containers)

	for _, container := range containers {
		go getContainerStats(container)
	}

	return nil
}

func calculateCPUPercentUnix(previousCPU, previousSystem uint64, v *types.StatsJSON) float64 {
	var (
		cpuPercent = 0.0
		cpuDelta = float64(v.CPUStats.CPUUsage.TotalUsage) - float64(previousCPU)
		systemDelta = float64(v.CPUStats.SystemUsage) - float64(previousSystem)
		onlineCPUs  = float64(v.CPUStats.OnlineCPUs)
	)

	if onlineCPUs == 0.0 {
		onlineCPUs = float64(len(v.CPUStats.CPUUsage.PercpuUsage))
	}
	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * onlineCPUs * 100.0
	}
	return cpuPercent
}

func getContainerStats(container types.Container) {
	time.Sleep(10 * time.Second)

	stats, err := cli.ContainerStats(ctx, container.ID, true)
	if err != nil {
		return
	}

	for {
		dec := json.NewDecoder(stats.Body)
		var v *types.StatsJSON

		if err := dec.Decode(&v); err != nil {
			if err == io.EOF {
				break
			}
			continue
		}

		var cpu = calculateCPUPercentUnix(v.PreCPUStats.CPUUsage.TotalUsage, v.PreCPUStats.SystemUsage, v)
		if cpu < 0.03 {
			down()
		} else if cpu > 40 {
			up()
		}

		fmt.Printf("%f %%\n", cpu)
	}
}