package cmd

import (
	"github.com/pojntfx/grpc-go-math/lib/math"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grpc-go-math-client",
	Short: "A client for `grpc-go-math-server`",
	Long:  "A simple CLI for the gRPC Go Math Server",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(math.MathClientCommand)
}
