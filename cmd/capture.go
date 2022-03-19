package cmd

import (
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
		captureOpt := &capture.Options{
			Devices: device,
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
}
