package capture

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"time"
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
	var handle *pcap.Handle
	timeout := time.Duration(opt.Timeout) * time.Second
	if opt.File == "" {
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
		// if file is specified
		handle, err = pcap.OpenOffline(opt.File)
	}
	if err != nil {
		return
	}

	packetSource = gopacket.NewPacketSource(
		handle,
		handle.LinkType(),
	)

	return
}
