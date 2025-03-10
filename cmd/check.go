/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the check command
var (
	targets  []string
	checkCmd = &cobra.Command{
		Use:   "check",
		Short: "Run connectivity checks for provided targets from configured datacenters.",
		Long:  `Runs connectivity checks for provided targets from configured datacenters.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Check connectivity for the following targets:")

			for _, v := range targets {
				fmt.Println(v)
			}

			jsonBody := []byte(fmt.Sprintf(`{"targets": %v}`, targets))
			bodyReader := bytes.NewReader(jsonBody)

			for _, v := range viper.GetStringSlice("datacenter") {
				fmt.Println(v)

				requestURL := fmt.Sprintf("%s/check", v)
				req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)

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
				fmt.Printf("Response: %s\n", responseBody)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringArrayVarP(&targets, "target", "t", []string{}, "Target ip addresses to check connectivity for")
	checkCmd.MarkFlagRequired("target")
}
