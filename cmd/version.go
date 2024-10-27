/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ECSbuilder",
	Long:  `Print the version number of ECSbuilder`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ECSbuilder version 0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
