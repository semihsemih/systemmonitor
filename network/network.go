package network

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type NetworkInfo struct {
	LocalIP string `json:"local_ip"`
}

var netInfo NetworkInfo

func localIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	hostAndPort := fmt.Sprintf("%v", localAddr)
	s := strings.Split(hostAndPort, ":")
	ip := s[0]

	return ip
}

func SystemNetworkInfo() NetworkInfo {
	netInfo.LocalIP = localIP()
	return netInfo
}
