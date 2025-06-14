package main

import (
	"fmt"
	"net"
)

// While studying Go, I came across this changelog email (https://groups.google.com/g/golang-announce/c/4t3lzH3I0eI), 
// This got me curious about what proxy bypassing means, what a zone ID is and how link-local IPv6 addresses work.
// I hadn’t explored low-level networking in Go before, so I wanted to understand this better.
//
// This program is the result of that curiosity. With help from ChatGPT, I wrote this to:
//   - Iterate over all network interfaces on the system.
//   - Skip interfaces that are down.
//   - For active interfaces, inspect their assigned addresses.
//   - Filter and print only IPv6 link-local unicast addresses (addresses that start with fe80::).
//   - Print them in the format `fe80::1%en0`, where `en0` is the interface name,
//     used to disambiguate between interfaces since link-local addresses are not globally unique.
//
// Writing and debugging this helped me understand net.Interfaces(), bitmask flags like FlagUp, type assertions
// in Go, and how IPv6 link-local addressing and zone identifiers are used in practice.
func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, iface := range interfaces {
		// Flags is a bitmask
		// Check if the interface is down by checking if its FlagUp flag is unset (== 0)
		// skip the interface if it is down or not configured
		if iface.Flags & net.FlagUp == 0 {
			continue
		}

		// fmt.Printf("Interface: %s (%v)\n", iface.Name, iface.Flags)

		addresses, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
			continue
		}


		for _, addr := range addresses {
			// Go type assertion that this addr of type net.Addr is actually an IP address of type net.IPAddr

			fmt.Printf("Address type: %T value: %v\n", addr, addr)
			ipnet, ok := addr.(*net.IPNet)

			// Skip the address, if it is
			// not an IP address
			// does not have an IP address value
			// not a valid IP address (either IPv4 or IPv6)
			// not a IPv6 address (or in other words, it is an IPv4 address)
			if !ok || ipnet.IP == nil || ipnet.IP.To16() == nil || ipnet.IP.To4() != nil {
				fmt.Printf("Not a valid candidate for an IPv6 link-local address\n\n")
				continue
			}

			// The IP address is an IPv6 address so now check if it's a link-local unicast address
			if ipnet.IP.IsLinkLocalUnicast() {
				fmt.Printf("Interface: %s\n", iface.Name)
				// Formatted to print an IPv6 link-local address as [fe80::1%eth0]
				// where fe80::1 is the IPv6 address and eth0 is the interface
				// In link-local addresses, an address can be used in multiple places so the interface is specified
				// First %s -> the IP address
				// % -> escape
				// % -> print a '%'
				// %s -> the interface name (also called zone ID) (eth0, wlan0, etc.)
				fmt.Printf("IPv6 Link-local address: [%s%%%s]\n\n", ipnet.IP.String(), iface.Name)
			} else {
				fmt.Printf("Not an IPv6 link-local address\n\n")
			}
		}

	}
}