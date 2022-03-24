package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yadavgulshan/packet-analyzer/capture"
)

var convertFile = &cobra.Command{
	Use:   "convert",
	Short: "convert pcap file to json",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		name, err := cmd.Flags().GetString("filename")
		if err != nil {
			return
		}
		errX := capture.CreateJson(name)
		if errX != nil {
			err = errX
			return
		}
		return
	},
}

func init() {
	rootcmd.AddCommand(convertFile)
	convertFile.Flags().StringP("filename", "f", "", "convert the existing pcap file to json file")
}
