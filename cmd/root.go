package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootcmd = &cobra.Command{
	Use:   "root",
	Short: "Root Command",
	Run:   run,
}

func run(cmd *cobra.Command, arg []string) {
	cmd.Usage()
}

func Execute() {
	if err := rootcmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
