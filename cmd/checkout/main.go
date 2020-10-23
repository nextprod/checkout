package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nextprod/checkout/pkg/sourceprovider"
	"github.com/spf13/cobra"
)

const nexWorkspace = "NEX_WORKSPACE"

var (
	cmd = &cobra.Command{
		Use: "checkout",
	}
	sshKey     string
	repository string
	ref        string
	runCmd     = &cobra.Command{
		Use:   "checkout",
		Short: "Checkout the repository",
		Long:  `The command checkouts repository according to provided command parameters`,
		Run:   run,
	}
)

func run(cmd *cobra.Command, args []string) {
	workspace := os.Getenv(nexWorkspace)
	if workspace == "" {
		panic(fmt.Errorf("checkout: workspace not found"))
	}
	git := sourceprovider.NewGitProvider()
	pkey, err := ioutil.ReadFile(sshKey)
	if err != nil {
		panic(err)
	}
	if err := git.Download(context.Background(), pkey, repository, ref, workspace); err != nil {
		panic(err)
	}
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(0)
	}
}

func init() {
	runCmd.PersistentFlags().StringVar(&sshKey, "ssh-key", "", "Private SSH key")
	runCmd.PersistentFlags().StringVar(&repository, "repository", "", "Repository name")
	runCmd.PersistentFlags().StringVar(&ref, "ref", "", "Reference name")

	cmd.AddCommand(runCmd)
	cmd.SetHelpCommand(&cobra.Command{
		DisableFlagsInUseLine: true,
		Hidden:                true,
	})
}
