/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile     string
	datacenters []string
	rootCmd     = &cobra.Command{
		Use:   "mario",
		Short: "Checking your networking plumbing",
		Long:  `A CLI tool to to send commands to a remote Mario instance and check connections`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mario.yaml)")
	rootCmd.PersistentFlags().StringArrayVarP(&datacenters, "datacentre", "d", []string{}, "Datacentres to check connectivity from")
	viper.BindPFlag("datacenter", rootCmd.PersistentFlags().Lookup("datacenter"))
	rootCmd.MarkPersistentFlagRequired("datacenter")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".mario" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".mario")
		viper.SafeWriteConfig()
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("mario")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
