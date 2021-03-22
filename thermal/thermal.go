package thermal

import (
	"strings"

	"github.com/codeskyblue/go-sh"
	"github.com/semihsemih/systemmonitor/utils"
)

type ThermalInfo struct {
	AOtherm  uint64
	CPUtherm uint64
	GPUtherm uint64
}

var thermInfo ThermalInfo

func DeviceThermalInfo() ThermalInfo {
	aoValueByte, _ := sh.Command("sudo", "cat", "/sys/devices/virtual/thermal/thermal_zone0/temp").Output()
	cpuValueByte, _ := sh.Command("sudo", "cat", "/sys/devices/virtual/thermal/thermal_zone1/temp").Output()
	gpuValueByte, _ := sh.Command("sudo", "cat", "/sys/devices/virtual/thermal/thermal_zone2/temp").Output()

	aoValueString := strings.Split(string(aoValueByte), "\n")[0]
	cpuValueString := strings.Split(string(cpuValueByte), "\n")[0]
	gpuValueString := strings.Split(string(gpuValueByte), "\n")[0]

	thermInfo.AOtherm = utils.StringToUint(aoValueString)
	thermInfo.CPUtherm = utils.StringToUint(cpuValueString)
	thermInfo.GPUtherm = utils.StringToUint(gpuValueString)

	return thermInfo

}
