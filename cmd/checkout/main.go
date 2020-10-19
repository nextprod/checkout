package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use: "checkout",
	}
	runCmd = &cobra.Command{
		Use:   "checkout",
		Short: "Checkout the repository",
		Long:  `The command checkouts repository according to provided command parameters`,
		Run:   run,
	}
)

func run(cmd *cobra.Command, args []string) {
	fmt.Println("OK")
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(0)
	}
}

func init() {
	cmd.AddCommand(runCmd)
	cmd.SetHelpCommand(&cobra.Command{
		DisableFlagsInUseLine: true,
		Hidden:                true,
	})
}
