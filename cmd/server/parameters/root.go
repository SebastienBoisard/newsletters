package parameters

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "newsletters-server",
	Short: "Backend",
	Long:  `Backend`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Usage()
		if err != nil {
			log.Printf("Error on usage [err: %v]", err)
		}
	},
}
