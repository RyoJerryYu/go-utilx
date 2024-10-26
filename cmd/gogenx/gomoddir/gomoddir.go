package gomoddir

import (
	"fmt"
	"os"

	"github.com/RyoJerryYu/go-utilx/pkg/utils/projectx"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func GoModDir() *cobra.Command {
	return &cobra.Command{
		Use:   "gomoddir",
		Short: "Print the directory of the go.mod file",
		Run: func(cmd *cobra.Command, args []string) {
			dir, err := projectx.GetGoModPath()
			if err != nil {
				glog.Fatalf("failed to get go.mod directory: %v", err)
				os.Exit(1)
			}

			fmt.Print(dir)
		},
	}
}
