package utils

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"time"
)

func IsFoundHost(host string, port uint16) bool {
	target := fmt.Sprintf("%s:%d", host, port)
	_, err := net.DialTimeout("tcp", target, 1*time.Second)
	if err != nil {
		fmt.Printf("%s %v\n", target, err)
		return false
	}
	return true
}

var PATTERN = regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)\.(\d+)$`)

func FindNeighbors(myHost string, myPort uint16, startIp uint8, endIp uint8, startPort uint16, endPort uint16) []string {
	address := fmt.Sprintf("%s:%d", myHost, myPort)
	fmt.Println(address)
	m := PATTERN.FindStringSubmatch(myHost)
	fmt.Println(m)
	if m == nil {
		return nil
	}
	prefixHost := fmt.Sprintf("%s.%s.%s", m[1], m[2], m[3])
	fmt.Println("Prefix Host:", prefixHost)
	lastIp, _ := strconv.Atoi(m[4])

	neighbors := make([]string, 0)
	for port := startPort; port <= endPort; port += 1 {
		for ip := startIp; ip <= endIp; ip += 1 {
			guessHost := fmt.Sprintf("%s%d", prefixHost, lastIp+int(ip))
			guessTarget := fmt.Sprintf("%s:%d", guessHost, port)
			if guessTarget != address && IsFoundHost(guessHost, port) {
				neighbors = append(neighbors, guessTarget)
			}
		}
	}
	return neighbors
}
