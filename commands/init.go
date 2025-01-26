package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/morgansundqvist/mh/config"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the configuration",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter root URL (e.g., http://localhost:8080/api): ")
		rootURL, _ := reader.ReadString('\n')
		rootURL = strings.TrimSpace(rootURL)

		if !config.IsValidURL(rootURL) {
			fmt.Println("Invalid URL. Please try again.")
			return
		}

		conf := config.Config{RootURL: rootURL}
		if err := config.SaveConfig(conf); err != nil {
			fmt.Println("Error saving configuration:", err)
			return
		}

		fmt.Println("Configuration saved successfully.")
	},
}
