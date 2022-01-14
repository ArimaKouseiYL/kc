/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package delete

import (
	"github.com/spf13/cobra"
	cmd2 "kc/cmd"
)

// deleteCmd represents the delete command
func NewCmdDelete() *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete resources by resources and names",
		Long: `Examples:
  
  # Delete pods  with same names "foo"
  kubectl delete pod foo
`,
		Run: func(cmd *cobra.Command, args []string) {

			resource, done := cmd2.Common(args)
			if done {
				return
			}

			if resource == "pods" || resource == "pod" || resource == "po" {

				podName := ""
				if len(args) > 1 {
					podName = args[1]
				}

				namespace, _ := cmd.Flags().GetString("namespace")

				deletePod(podName, namespace)
			}
		},
	}

	deleteCmd.Flags().StringP("namespace", "n", "", "specify pod namespace")

	return deleteCmd
}
