package capture

import (
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type PacketSourceOpt struct {
	Device      string
	BPFFilter   string
	Promiscous  bool
	SnapShotLen int32
	Timeout     int
	File        string
}

func CreatePacketSource(opt *PacketSourceOpt) (
	packetSource *gopacket.PacketSource,
	err error,
) {
	println("Create Packet source")
	var handle *pcap.Handle
	timeout := time.Duration(opt.Timeout) * time.Second
	if opt.File == "" {
		println("Starting live mode")
		//OpenLive opens a device and returns a *Handle. It takes as arguments the name
		//of the device ("eth0"), the maximum size to read for each packet (snaplen),
		//whether to put the interface in promiscuous mode, and a timeout. Warning: this
		//function supports only microsecond timestamps. For nanosecond resolution use
		//an InactiveHandle.
		handle, err = pcap.OpenLive(
			opt.Device,
			opt.SnapShotLen,
			opt.Promiscous,
			timeout,
		)
	} else {
		// This section won't get executed at all!
		// why?
		// Because I didn't specified the file object when calling this function.
		// So by default this function will open the live mode.
		// Then this live mode will be captured inside a file.
		println("Trying to open file")
		// if file is specified
		handle, err = pcap.OpenOffline(opt.File)
	}
	if err != nil {
		return
	}
	if opt.BPFFilter != "" {
		err = handle.SetBPFFilter(opt.BPFFilter)
		if err != nil {
			return
		}
	}

	packetSource = gopacket.NewPacketSource(
		handle,
		handle.LinkType(),
	)

	return
}
