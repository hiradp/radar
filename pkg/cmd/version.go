package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/hiradp/radar/internal/build"
)

var versionCmd = &cobra.Command{
	Use:    "version",
	Hidden: true,
	Short:  "Print version and exit",
	Args:   cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		info := build.Info{}
		fmt.Println(info)
	},
}
