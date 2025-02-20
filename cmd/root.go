package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Sun",
	Short: "Sun is a forecast CLI application",
	Long: `Sun is a command line interface (CLI) application that provides weather forecasts.
It allows users to get current weather conditions, forecasts for the upcoming days, and other weather-related information.
With Sun, you can easily check the weather for any location directly from your terminal, making it a handy tool for developers and terminal enthusiasts.`,
	Run: func(cmd *cobra.Command, args []string) {
		sunIcon := `
  _____ _   _ _   _ 
 / ____| | | | \ | |
| (___ | | | |  \| |
 \___ \| | | | . \ |
 ____) | |_| | |\  |
|_____/ \___/|_| \_|
		`
		fmt.Println(sunIcon)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
