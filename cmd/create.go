/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var (
	special bool
	length  int
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates password for a given application name",
	Long: `Creates a password for a given application name. 
	Provides the ability to include special characters and set the length of the password.`,
	Run: func(cmd *cobra.Command, args []string) {
		// need to validate that the user has provided a name

		fmt.Printf("Creating a password for %s\n", args[0])
		chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		if special {
			chars += "!@#$%^&*()_+"
		}
		password := make([]byte, length)
		for i := range password {
			password[i] = chars[rand.Intn(len(chars))]
		}
		Passwords[args[0]] = string(password)
		fmt.Println(Passwords)
	},
}

func init() {

	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.
	createCmd.Flags().BoolVarP(&special, "special", "s", false, "Include special characters")
	createCmd.Flags().IntVarP(&length, "length", "l", 8, "Length of the password")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
