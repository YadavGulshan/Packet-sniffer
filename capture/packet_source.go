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
