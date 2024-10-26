package main

import (
	"os"

	"github.com/RyoJerryYu/go-utilx/cmd/gogenx/gomoddir"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use:   "gogenx",
		Short: "Command line tool helper for go generate",
	}
	cmd.AddCommand(
		gomoddir.GoModDir(),
	)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
