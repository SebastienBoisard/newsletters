package main

import (
	"fmt"
	"github.com/SebastienBoisard/newsletters/cmd/server/parameters"
	"os"
)

func main() {

	if err := parameters.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
