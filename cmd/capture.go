package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yadavgulshan/packet-analyzer/capture"
)

var captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Capture packets",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		device, err := cmd.Flags().GetStringArray("devices")
		if err != nil {
			return err
		}

		file, err := cmd.Flags().GetString("file")
		if err != nil {
			return
		}
		filter, err := cmd.Flags().GetString("filter")
		if err != nil {
			return
		}
		maxPackets, err := cmd.Flags().GetInt("max_packets")
		if err != nil {
			return
		}
		if maxPackets < 0 {
			err = fmt.Errorf("max packets cannot be <= 0")
			return
		}
		captureOpt := &capture.Options{
			Devices:    device,
			File:       file,
			BPFFilter:  filter,
			MaxPackets: maxPackets,
		}
		err = captureOpt.Validate()
		if err != nil {
			return
		}
		return
	}}

func init() {
	rootcmd.AddCommand(captureCmd)
	captureCmd.Flags().StringArrayP("devices", "d", []string{}, "Devices to capture")
	captureCmd.Flags().StringP("file", "f", "", "Pcap file")
	captureCmd.Flags().StringP("filter", "F", "", "BPF filter")
	captureCmd.Flags().IntP("max_packets", "N", 0, "Maximum number of packets to capture")
}
