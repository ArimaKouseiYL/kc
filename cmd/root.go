/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"kc/cmd/create"
	delete2 "kc/cmd/delete"
	"kc/cmd/get"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kc",
	Short: "kc controls the Kubernetes cluster manager.",
	Long: `kc controls the Kubernetes cluster manager.

Basic Commands (Beginner):
  create         Create a resource from a file or from stdin.
  get            Display one or many resources
  edit           Edit a resource on the server
  delete         Delete resources by filenames, stdin, resources and names, or by resources and label selector`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(create.NewCmdCreate())
	rootCmd.AddCommand(delete2.NewCmdDelete())
	rootCmd.AddCommand(get.NewCmdGet())
}
