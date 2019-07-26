package cmd

import (
	"github.com/spf13/cobra"
)

func newStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start HTTP Server",
		Run: func(_ *cobra.Command, _ []string) {
			initDependency()
			httpServer.StartHTTPServer()
		},
	}
}
