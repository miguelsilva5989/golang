package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  string
	State string
}

// ScanPort scans a port for a specific protocol, hostname
func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{
		Port: protocol + "/" + strconv.Itoa(port),
	}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		return result
	}
	defer conn.Close()

	result.State = "Open"
	return result
}

// InitialScan scan first 1024 ports
func InitialScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
		results = append(results, ScanPort("udp", hostname, i))
	}

	return results
}
