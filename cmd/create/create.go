/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package create

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create resourceType",
		DisableFlagsInUseLine: true,
		Short:                 "Create a resource from stdin.",
		Long: `Create a resource from from stdin.

Examples:
  # Create a pod using only special podName and image. 
  kubectl create pod pod-test -i nginx:latest
  
  # Create a pod using special podName、image and namespace.
  kubectl create pod pod-test -i nginx:latest -n example
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

				namespace, err := cmd.Flags().GetString("namespace")
				if err != nil {
					return
				}

				image, err := cmd.Flags().GetString("image")
				if err != nil {
					return
				}

				port, err := cmd.Flags().GetInt32("port")
				if err != nil {
					return
				}

				createPod(podName, namespace, image, port)

			}

		},
	}

	cmd.Flags().StringP("namespace", "n", "", "specify pod namespace")
	cmd.Flags().StringP("image", "i", "", "specify pod image")
	cmd.Flags().Int32P("port", "p", 80, "specify pod port")

	return cmd
}
