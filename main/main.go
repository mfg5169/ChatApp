package main

import (
	"fmt"
	"log"
	"net"
	"github.com/vishvananda/netlink"
)

func main() {
	// Define the virtual IP and network interface
	vip := "192.168.1.100/32" 
	ifaceName := "eth0"       

	// Parse the virtual IP
	ip, ipNet, err := net.ParseCIDR(vip)
	if err != nil {
		log.Fatalf("Invalid IP address: %v", err)
	}
	ipNet.IP = ip

	// Get the network interface
	link, err := netlink.LinkByName(ifaceName)
	if err != nil {
		log.Fatalf("Failed to find interface %s: %v", ifaceName, err)
	}

	// Add the VIP to the interface
	addr := &netlink.Addr{IPNet: ipNet, Label: ""}
	if err := netlink.AddrAdd(link, addr); err != nil {
		log.Fatalf("Failed to add VIP %s to interface %s: %v", vip, ifaceName, err)
	}

	fmt.Printf("Successfully assigned VIP %s to interface %s\n", vip, ifaceName)
}
