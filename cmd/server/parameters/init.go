package parameters

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the server",
	Long:  `Initialize the server (clean and fill the database; ...)`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// userPassword := args[0]
		// docshare_backend.InitBackend(userPassword)
	},
}
