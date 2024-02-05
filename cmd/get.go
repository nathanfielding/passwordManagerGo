/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Allows you to retrive a single or all passwords",
	Long:  "Allows you to retrive a single or all passwords",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		passwords, err := loadPasswords("") // need to figure out where to store file
		if err != nil {
			return err
		}

		if args[0] == "all" {
			if len(passwords) < 1 {
				return fmt.Errorf("no passwords exist")
			}
			for k, v := range passwords {
				fmt.Printf("%s: %s\n", k, v)
			}
			return nil
		}
		app := args[0]
		v, ok := passwords[app]
		if !ok {
			return fmt.Errorf("password does exist for %s", app)
		}
		fmt.Printf("%s: %s", app, v)
		return nil
	},
}

func loadPasswords(filename string) (map[string]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var passwords map[string]string
	if err := json.Unmarshal(data, &passwords); err != nil {
		return nil, err
	}
	return passwords, nil
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
