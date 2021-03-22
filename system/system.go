package system

import (
	"github.com/semihsemih/systemmonitor/cpu"
	"github.com/semihsemih/systemmonitor/disk"
	h "github.com/semihsemih/systemmonitor/host"
	"github.com/semihsemih/systemmonitor/memory"
	"github.com/semihsemih/systemmonitor/network"
	"github.com/semihsemih/systemmonitor/thermal"
	"github.com/shirou/gopsutil/v3/host"
)

type SystemInfo struct {
	HostInfo *host.InfoStat      `json:"host_info"`
	Memory   memory.MemoryInfo   `json:"memory"`
	Network  network.NetworkInfo `json:"network"`
	Disk     disk.DiskInfo       `json:"disk"`
	Thermal  thermal.ThermalInfo `json:"thermal"`
	CPU      cpu.CPUInfo         `json:"cpu"`
}

var sysInfo SystemInfo

func DeviceSystemInfo() SystemInfo {
	sysInfo.Memory = memory.DeviceMemoryInfo()
	sysInfo.HostInfo = h.SystemInfo()
	sysInfo.Network = network.SystemNetworkInfo()
	sysInfo.Disk = disk.DiskUsageInfo()
	sysInfo.Thermal = thermal.DeviceThermalInfo()
	sysInfo.CPU = cpu.DeviceCPUInfo()

	return sysInfo
}
