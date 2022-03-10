package main

import "fmt"
import "strings"

type IPAddr [4]byte

func (p IPAddr) String() string {
	return strings.Trim(strings.Replace(fmt.Sprint(p[:]), " ", ".", -1), "[]")
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
