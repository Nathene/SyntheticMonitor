package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "synthetic-monitor",
	Short: "A tool for monitoring the health of services and endpoints",
	Long:  "Synthetic Monitor is a CLI tool designed to monitor the health and performance of various services.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() // Default action: show help
	},
}

// Execute initializes the root command and executes it.
func Execute(ctx context.Context) {
	cobra.OnInitialize(initConfig)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// initConfig reads in the configuration file and environment variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use the specified config file
		viper.SetConfigFile(cfgFile)
	} else {
		// Default to looking for a config file in the current directory
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	// Read in environment variables that match
	viper.AutomaticEnv()

	// Read the configuration file
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Printf("Error reading config file: %v\n", err)
	}
}

func init() {
	// Add a flag for specifying the config file
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")

	// Bind Viper to the config flag
	_ = viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}
