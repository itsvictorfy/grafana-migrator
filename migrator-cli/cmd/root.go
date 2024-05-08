/*
Copyright © 2024 Victor Fun-Young victor@funyoung.org
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	source  string
	target  string
	token   string
	verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grafana-migrator",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Check if the version flag is set
	if versionFlagSet() {
		// If the version flag is set, print the version and exit
		fmt.Printf("Grafana Migrator Version: 1.0.2")
		return
	}

	// If the version flag is not set, execute the root command
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// versionFlagSet checks if the version flag is set
func versionFlagSet() bool {
	versionFlag, _ := rootCmd.Flags().GetBool("version")
	return versionFlag
}

func init() {
	// Flags for source, target, and token
	rootCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "Source Grafana instance")
	rootCmd.PersistentFlags().StringVarP(&target, "target", "t", "", "Target Grafana instance")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "k", "", "User token")

	// Verbose flag
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	// Help flag
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Help for "+rootCmd.Name())

	// Version flag
	rootCmd.PersistentFlags().BoolP("version", "V", false, "Print the version number")
}
