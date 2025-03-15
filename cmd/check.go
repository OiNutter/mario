/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CheckRequest struct {
	Targets []string `json:"targets"`
}
type CheckResult struct {
	Status string `json:"status" example:"success" description:"Status of the connection check."`
	Error  string `json:"error,omitempty" example:"connection refused" description:"Error message if the connection check failed."`
}

type CheckResponse map[string]CheckResult

// checkCmd represents the check command
var (
	targets  []string
	checkCmd = &cobra.Command{
		Use:   "check",
		Short: "Run connectivity checks for provided targets from configured datacenters.",
		Long:  `Runs connectivity checks for provided targets from configured datacenters.`,
		Run: func(cmd *cobra.Command, args []string) {
			jsonBody, err := json.Marshal(&CheckRequest{
				Targets: targets,
			})
			if err != nil {
				fmt.Printf("Error marshalling request: %v\n", err)
			}
			bodyReader := bytes.NewBuffer(jsonBody)

			for _, v := range viper.GetStringSlice("datacenter") {

				requestURL := fmt.Sprintf("%s/check", v)
				req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)

				req.Header.Set("Content-Type", "application/json; charset=UTF-8")

				if err != nil {
					fmt.Printf("Error creating request: %v\n", err)
					return
				}

				res, err := http.DefaultClient.Do(req)
				if err != nil {
					fmt.Printf("Error sending request: %v\n", err)
					return
				}

				responseBody, err := io.ReadAll(res.Body)
				if err != nil {
					fmt.Printf("Error reading response: %v\n", err)
					return
				}

				var response CheckResponse
				json.Unmarshal(responseBody, &response)

				var anyFailed bool
				for _, v := range response {
					if v.Status != "success" {
						anyFailed = true
						break
					}
				}

				var c *color.Color
				if anyFailed {
					c = color.New(color.FgRed)
				} else {
					c = color.New(color.FgGreen)
				}

				c.Println("\n####################")
				c.Println("Datacentre:", v)

				for k, v := range response {
					var d *color.Color
					if v.Status == "success" {
						d = color.New(color.FgGreen)
					} else {
						d = color.New(color.FgRed)
					}

					fmt.Println("---")
					d.Println("Target:", k)
					d.Println("Status:", v.Status)
					if v.Error != "" {
						d.Println("Error:", v.Error)
					}
				}
				c.Println("####################")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringArrayVarP(&targets, "target", "t", []string{}, "Target ip addresses to check connectivity for")
	checkCmd.MarkFlagRequired("target")
}
