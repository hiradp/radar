package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/hiradp/radar/pkg/radar"
)

var scanCmd = &cobra.Command{
	Use:   "scan <host>",
	Short: "Scan a host for relevant SSL information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		h := args[0]
		f, err := cmd.Flags().GetString("format")
		if err != nil {
			log.Println("ERROR failed to get flags", err)
			fmt.Println("Unexpected error occurred, please check the log file for more information")
			return
		}

		format := strings.ToLower(f)


		host, err := radar.Scan(h)
		if err != nil {
			log.Fatalln(err)
		}

		if format == "plain" {
			fmt.Println(host)
		} else if format == "html" {
			fmt.Println(host.HTML())
		} else {
			fmt.Printf("Unknown output format of %s - try using 'html' or 'plain'\n")
		}
	},
}

func init() {
	scanCmd.Flags().StringP("format", "f", "plain", "The output format. Either \"plain\" or \"html\"")
}
