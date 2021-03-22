package disk

import (
	"strings"

	"github.com/codeskyblue/go-sh"
	"github.com/semihsemih/systemmonitor/utils"
)

type DiskInfo struct {
	Size      uint64 `json:"size"`
	Used      uint64 `json:"used"`
	Available uint64 `json:"available"`
}

var info DiskInfo

func DiskUsageInfo() DiskInfo {
	output, _ := sh.Command("df", "-l").Command("sed", "-n", "2p").Output()
	outputString := string(output)
	replacedString := strings.Replace(outputString, "  ", " ", -1)
	splittedString := strings.Split(replacedString, " ")

	info.Size = utils.StringToUint(splittedString[1])
	info.Used = utils.StringToUint(splittedString[2])
	info.Available = utils.StringToUint(splittedString[3])

	return info
}
