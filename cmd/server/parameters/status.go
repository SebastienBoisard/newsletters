package parameters

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Retrieve the current server status",
	Long:  `Retrieve the current server status`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("server status\n")
	},
}
