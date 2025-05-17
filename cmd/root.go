package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func handler(data []string) {
	operation := data[0]
	switch operation {
	case "analyze":
		display(data[1])
	case "rename":
		renameFile(data[1], data[2])
	}
}

var rootCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze is used for analyzing a text file",
	Long:  "A nice little text analyzer for friendly text analysis that you could simply achieve using bash commands such as wc",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handler(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
