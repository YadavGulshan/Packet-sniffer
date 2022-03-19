package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yadavgulshan/packet-analyzer/capture"
)

var listDevicesCmd = &cobra.Command{
	Use:   "list_devices",
	Short: "list network devices",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := capture.ListDevices()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootcmd.AddCommand(listDevicesCmd)
}
