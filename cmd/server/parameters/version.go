package parameters

import (
	"fmt"
	"github.com/SebastienBoisard/newsletters/internal/server"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the server",
	Long:  `All software has versions. This is server's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v%s\n", server.ProjectName, server.ProjectVersion)
	},
}
