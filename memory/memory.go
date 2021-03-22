package memory

import (
	"github.com/shirou/gopsutil/v3/mem"
)

type MemoryInfo struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
}

var memInfo MemoryInfo

func systemUsedMemory() uint64 {
	v, _ := mem.VirtualMemory()
	return v.Used
}

func totalMemory() uint64 {
	v, _ := mem.VirtualMemory()
	return v.Total
}

func DeviceMemoryInfo() MemoryInfo {
	memInfo.Total = totalMemory()
	memInfo.Used = systemUsedMemory()

	return memInfo
}
