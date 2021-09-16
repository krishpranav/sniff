package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

type sliceVal []string

func (s sliceVal) String() string {
	var str string
	for _, i := range s {
		str += fmt.Sprintf("%s\n", i)
	}
	return str
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("[!] Couldn't create file: %s\n", err.Error())
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintf(w, line)
	}

	return w.Flush()
}

func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func IsIPv6CIDR(cidr string) bool {
	ip, _, _ := net.ParseCIDR(cidr)
	return ip != nil && ip.To4() == nil
}

func CIDRToIP(cidr string) []string {
	ip, ipnet, err := net.ParseCIDR(cidr)

	if err != nil {
		log.Fatalf("[!] Failed")
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incrementIP(ip) {
		ips = append(ips, ip.String())
	}

	if len(ips) >= 3 {
		return ips[1 : len(ips)-1]
	} else {
		return ips
	}
}

func getIP(ipdomain string) []net.IP {
	ip, err := net.LookupIP(ipdomain)
	if err != nil {
		log.Fatalf("[!] Error loooking up domain: %s\n", err.Error())
	}

	return ip
}

func main() {
	var (
		target = flag.String("t", "", "Domain or ip address (Required!!)")
		print  = flag.Bool("p", false, "Print Result to console")
	)

	flag.Parse()
}
