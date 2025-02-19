package cli

import (
	"github.com/spf13/cobra"
	"log"
)

var cmd = &cobra.Command{
	Use:   "algo",
	Short: "Help to solve algorithm problems",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	cmd.AddCommand(addCmd)
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command: %v\n", err)
	}
}
