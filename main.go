package main

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

// GetOutboundIP Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// Task for ticker
func Task() {
	ip := GetOutboundIP()
	fmt.Println(ip)
	cmd := fmt.Sprintf("echo %v > ip", ip)
	output, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(output))

	output, err = exec.Command("bash", "-c", "scp ip $s88h/").Output()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(output))
}

func main() {
	Task()

	ticker := time.NewTicker(time.Hour * 24)
	for {
		<-ticker.C
		go Task()
	}
}
