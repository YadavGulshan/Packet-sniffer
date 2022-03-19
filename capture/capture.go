package capture

import (
	"fmt"
	"github.com/google/gopacket/pcap"
)

func ListDevices() (err error) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return
	}
	fmt.Println("Found Devices")
	for _, device := range devices {
		fmt.Println("--------------------------------")
		fmt.Printf("Device: %s\n", device.Name)
		fmt.Printf("Description: %s\n", device.Description)
		fmt.Printf("Flags: %d\n", device.Flags)
		fmt.Printf("Address: %s\n", device.Addresses)
	}
	return
}
