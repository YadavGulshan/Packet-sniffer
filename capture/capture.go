package capture

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type Options struct {
	// List of devices user want's to sniff.
	Devices []string
	//	Berkeley packet filters
	BPFFilter   string
	Promiscous  bool
	SnapshotLen int32
	Timeout     int
	File        string
	MaxPackets  int
}

const (
	defaultTimeOut     = 30
	defaultSnapshotLen = 1024
)

// Validate function
func (opt *Options) Validate() (err error) {
	if len(opt.Devices) < 1 {
		err = fmt.Errorf("provide a device to sniff")
		return
	}

	// Check if all the devices mentioned exists in the current device.
	for _, device := range opt.Devices {
		ok, errX := IsDeviceExists(device)
		err = errX
		if !ok {
			err = fmt.Errorf("device Name with %s does not exists", device)
			return
		}
	}
	if opt.SnapshotLen == 0 {
		opt.SnapshotLen = defaultSnapshotLen
	}
	if opt.Timeout == 0 {
		opt.Timeout = defaultTimeOut
	}
	return
}

func IsDeviceExists(name string) (ok bool, err error) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return false, err
	}
	for _, device := range devices {
		if name == device.Name {
			println("device exists!")
			ok = true
			return ok, nil
		}
	}
	return false, nil
}

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

func CreatePacketCaptureSource(opt *Options) (packetSources []*gopacket.PacketSource, err error) {
	packetSources = []*gopacket.PacketSource{}
	for _, device := range opt.Devices {
		pktSourceOpt := &PacketSourceOpt{
			device,
			opt.BPFFilter,
			opt.Promiscous,
			opt.SnapshotLen,
			opt.Timeout,
			opt.File,
		}

		packetSource, errX := CreatePacketSource(pktSourceOpt)

		if errX != nil {
			err = errX
			return
		}
		packetSources = append(packetSources, packetSource)

	}
	return
}
