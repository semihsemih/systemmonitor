package cpu

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type CPUInfo struct {
	Percent float64 `json:"percent"`
}

var cInfo CPUInfo

func cpuPercentUsage() float64 {
	percent, _ := cpu.Percent(time.Second*3, false)
	return percent[0]
}

func DeviceCPUInfo() CPUInfo {
	cInfo.Percent = cpuPercentUsage()
	return cInfo
}
