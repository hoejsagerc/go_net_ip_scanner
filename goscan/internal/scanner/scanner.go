package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type ScanResult struct {
	IP       string
	Port     int
	Open     bool
	Hostname string
}

func StartScan(startIp, endIp string, timeout time.Duration, ports []int) ([]ScanResult, error) {
	ips, err := getIPRange(startIp, endIp)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	resultsChan := make(chan ScanResult, len(ips)*len(ports))

	var portWg sync.WaitGroup

	for _, ip := range ips {
		for _, port := range ports {
			portWg.Add(1)
			wg.Add(1)
			go func(ip string, port int) {
				defer wg.Done()
				resultsChan <- scanPort(ip, port, timeout)
				portWg.Done()
			}(ip, port)
		}
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var results []ScanResult
	for result := range resultsChan {
		results = append(results, result)
	}

	return results, nil
}

func getIPRange(startIP, endIP string) ([]string, error) {
	var ips []string
	start := net.ParseIP(startIP)
	end := net.ParseIP(endIP)

	for ip := start; !ip.Equal(end); ip = nextIP(ip) {
		ips = append(ips, ip.String())
	}
	ips = append(ips, end.String())

	return ips, nil
}

func nextIP(ip net.IP) net.IP {
	ip = ip.To4()
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
	return ip
}

func scanPort(ip string, port int, timeout time.Duration) ScanResult {
	fmt.Printf("Scanning %s:%d\n", ip, port)
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return ScanResult{IP: ip, Port: port, Open: false}
	}
	conn.Close()

	hostnames, err := net.LookupAddr(ip)
	if err != nil || len(hostnames) == 0 {
		return ScanResult{IP: ip, Port: port, Open: true, Hostname: ""}
	}

	return ScanResult{IP: ip, Port: port, Open: true, Hostname: hostnames[0]}
}
