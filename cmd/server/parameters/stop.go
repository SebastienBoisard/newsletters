package parameters

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the server",
	Long:  `Stop the server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("stop server\n")
	},
}
