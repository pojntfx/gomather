package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/z0mbie42/rz-go/v2"
	"gitlab.com/z0mbie42/rz-go/v2/log"
)

var RootCmd = &cobra.Command{
	Version: "1.0.0",
	Use:     "gomather",
	Short:   "Simple Go gRPC microservice that does math.",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal("Could not start root command", rz.Err(err))
	}
}
