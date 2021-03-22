package main

import (
	"encoding/json"
	"fmt"

	"github.com/semihsemih/systemmonitor/system"
)

func main() {
	info := system.DeviceSystemInfo()
	json, _ := json.Marshal(info)
	fmt.Println(string(json))
}
