package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/privateerproj/privateer-raid-example/armory"
)

var (
	// debugCmd represents the base command when called without any subcommands
	debugCmd = &cobra.Command{
		Use:   "debug",
		Short: "Run the Raid in debug mode",
		Run: func(cmd *cobra.Command, args []string) {
			err := Vessel.Mobilize(&armory.Armory, nil) // Replace nil with a slice of your required var names
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	runCmd.AddCommand(debugCmd) // This enables the debug command for use while working on your Raid
}
