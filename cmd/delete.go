/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a password for a given application name",
	Long:  `Deletes a password for a given application name.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Deleting password for %s\n", args[0])
		app := args[0]
		_, ok := Passwords[app]
		if !ok {
			return fmt.Errorf("password does not exist for %s", app)
		}
		fmt.Println("Sucessfully deleted password")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
