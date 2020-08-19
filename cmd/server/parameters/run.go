package parameters

import (
	"errors"
	"fmt"
	"github.com/SebastienBoisard/newsletters/internal/server"
	"github.com/spf13/cobra"
	"strconv"
)

var releaseMode bool

func init() {
	RootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolVarP(&releaseMode, "release_mode", "r", false, "run the server in release mode")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Long:  `Run the server`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires the http port number of the demo")
		}
		_, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("%q must be a number.\n", args[0])
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		portNumber, _ := strconv.Atoi(args[0])
		fmt.Println("Run server on port ", portNumber)
		server.Run(portNumber, releaseMode)
	},
}
