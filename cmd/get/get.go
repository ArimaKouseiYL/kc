/*
Copyright © 2022 <Mark>
*/
package get

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cmd2 "kc/cmd"
	"kc/config"
	"kc/k8s"
)

func NewCmdGet() *cobra.Command {
	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Display one or many resources",
		Long: `Use "kc api-resources" for a complete list of supported resources.

Examples:
  # List all pods in ps output format.
  kubectl get pods
  
  # List a single pod in JSON output format.
  kubectl get -o json pod web-pod-13je7 

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

				output, err := cmd.Flags().GetString("output")
				if err != nil {
					return
				}

				all, err := cmd.Flags().GetBool("all")
				if err != nil {
					return
				}
				getPod(podName, namespace, output, all)
			}

		},
	}

	getCmd.Flags().StringP("namespace", "n", "", "specify pod namespace")
	getCmd.Flags().StringP("output", "o", "", "specify api-resources output format：json|yaml")
	getCmd.Flags().BoolP("all", "A", false, "show all namespace pod resources ")
	return getCmd
}

// getCmd represents the get command

func PodListOutPut(namespace string) bool {
	podList, err := config.K8sClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return true
	}
	if len(podList.Items) == 0 {
		fmt.Println("No resources found in default namespace.")
	} else {
		k8s.StandardOutPut(podList)
	}
	return false
}
