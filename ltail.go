package main

import (
	"os"

	"github.com/megaproaktiv/ltail/cmd"
)

func main() {
	if err := cmd.LtailCommand.Execute(); err != nil {
		os.Exit(0)
	}
}
