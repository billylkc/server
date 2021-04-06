package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Host server.",
	Long:  `Host server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Hosting on port - %d \n", port)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(hostCmd)
	hostCmd.Flags().IntVarP(&port, "port", "p", 3000, "port")
}
