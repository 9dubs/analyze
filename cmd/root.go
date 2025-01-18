package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze is used for analyzing a text file",
	Long:  "A nice little text analyzer for friendly text analysis that you could simply achieve using bash commands such as wc",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(args[0])
		check(err)
		content := http.DetectContentType(data)
		fmt.Println("The file content type is: " + content)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
