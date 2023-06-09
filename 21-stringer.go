package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (t IPAddr) String() (result string) {
	result = fmt.Sprintf("%d.%d.%d.%d", t[0], t[1], t[2], t[3])
	return
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	fmt.Println("=====================")
	for name, ip := range hosts {
		fmt.Print(name)
		fmt.Print("  ")
		fmt.Print(ip.String())
		fmt.Print("\n")
	}
}
