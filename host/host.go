package host

import (
	"github.com/shirou/gopsutil/v3/host"
)

func SystemInfo() *host.InfoStat {
	h, _ := host.Info()
	return h
}

