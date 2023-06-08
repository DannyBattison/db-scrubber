/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"danny.gg/db-scrubber/internal/cfg"
	"danny.gg/db-scrubber/internal/db/mysql"
	"github.com/spf13/cobra"
)

// scrubCmd represents the scrub command
var scrubCmd = &cobra.Command{
	Use:   "scrub [config file]",
	Short: "Scrubs a database",
	Long:  "A tool which scrubs a database with a given configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := cfg.Load(args[0])
		if err != nil {
			panic(err)
		}

		err = mysql.Scrub(config)

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(scrubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scrubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scrubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
