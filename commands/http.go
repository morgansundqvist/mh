package commands

import (
	"fmt"

	"github.com/morgansundqvist/mh/config"
	"github.com/morgansundqvist/mh/httpclient"
	"github.com/spf13/cobra"
)

var GetCmd = createHTTPCommand("get", "GET", "Get data from the API")
var PostCmd = createHTTPCommand("post", "POST", "Post data to the API")
var PutCmd = createHTTPCommand("put", "PUT", "Update data on the API")
var DeleteCmd = createHTTPCommand("delete", "DELETE", "Delete data from the API")
var PatchCmd = createHTTPCommand("patch", "PATCH", "Patch data from the API")

func createHTTPCommand(name, method, description string) *cobra.Command {
	var output bool

	returnCmd := &cobra.Command{
		Use:   name + " [url] [key=value...]",
		Short: description,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			conf, err := config.LoadConfig()
			if err != nil {
				fmt.Println("Error loading configuration:", err)
				return
			}

			urlPath := args[0]
			params := args[1:]

			req, err := httpclient.CreateRequest(method, conf.RootURL, urlPath, params)
			if err != nil {
				fmt.Println("Error creating request:", err)
				return
			}

			if err := httpclient.ExecuteRequest(req, output); err != nil {
				fmt.Println("Error executing request:", err)
			}
		},
	}

	returnCmd.Flags().BoolVarP(&output, "output", "o", false, "Output response status")

	return returnCmd
}
