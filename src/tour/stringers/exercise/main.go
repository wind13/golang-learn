package main

import (
	"fmt"
	"strings"
	"strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ipAddr IPAddr) String() string {
	ips := make([]string, 4)
	for i, ip := range ipAddr {
		// fmt.Println(i, ip)
		// fmt.Printf("%T", ip)
		ips[i] = strconv.Itoa(int(ip))
	}
	// fmt.Println(ips)
	return fmt.Sprintf("%v", strings.Join(ips, "."))
	// return fmt.Sprint(strings.Join([]string{"a", "b", "c"}, "."))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
