/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Set up your Mario config",
	Long:  `Takes you through setting up your Mario configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("We'll now take you through setting up your Mario configuration file.")
		fmt.Println("---------------------------------")

		fmt.Println("Please enter the datacenters you'd like to check connectivity from. Enter a blank line to finish")
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		datacenters := []string{}
		text = strings.Replace(text, "\n", "", -1)

		for text != "" {
			datacenters = append(datacenters, text)
			fmt.Print("-> ")
			text, _ = reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
		}

		viper.Set("datacenter", datacenters)

		viper.WriteConfig()

	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
