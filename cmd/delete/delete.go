/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package delete

import (
	"fmt"
	"github.com/spf13/cobra"
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

			resource := ""
			if len(args) == 0 {
				fmt.Println("You must specify the type of resource to get. eg：kc get po")
				return
			}
			resource = args[0]

			if resource == "pods" || resource == "pod" || resource == "po" {

				podName := ""
				if len(args) > 1 {
					podName = args[1]
				}

				namespace, _ := cmd.Flags().GetString("namespace")

				if namespace == "" {
					namespace = "default"
				}

				deletePod(podName, namespace)
			}
		},
	}

	deleteCmd.Flags().StringP("namespace", "n", "", "specify pod namespace")

	return deleteCmd
}
