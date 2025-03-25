/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Category name:", category)
		fmt.Println("Category exists:", exists)
		fmt.Println("Category ID:", id)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking if category exists")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Category command completed")
	},
}

var category string
var exists bool
var id int16

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "Y", "Name of the category")
	categoryCmd.PersistentFlags().BoolVarP(&exists, "exists", "e", false, "Check if the category exists")
	categoryCmd.PersistentFlags().Int16VarP(&id, "id", "i", 0, "ID of the category")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
