/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/hoejsagerc/go_net_ip_scanner/goscan/cmd/scan"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goscan",
	Short: "A tcp network scanner written in go",
	Long: `A network scanner for scanning tcp ports and locating devices on either
	local network or remote networks. The scanner can be used to scan a single or a range of
	ip addresses and ports.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPalletes() {
	rootCmd.AddCommand(scan.ScanCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubCommandPalletes()
}
