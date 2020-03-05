package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/bloom42/libs/rz-go"
	"gitlab.com/bloom42/libs/rz-go/log"
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
