package get

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kc/config"
	"kc/k8s"
)

func getPod(name, namespace, output string, all bool) {

	if namespace == "" {
		if all {
			for _, namespace := range k8s.GetNamespaceNames() {
				if PodListOutPut(namespace.Namespace) {
					return
				}
			}
		} else {
			if PodListOutPut("default") {
				return
			}
		}
	} else {
		if name != "" {
			pod, _ := config.K8sClient.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})

			var pList []v1.Pod
			pods := append(pList, *pod)

			list := &v1.PodList{
				Items: pods,
			}
			k8s.StandardOutPut(list)

		} else {
			podList, _ := config.K8sClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

			if len(podList.Items) == 0 {
				fmt.Printf("No resources found in %s namespace.\n", namespace)
			} else {
				k8s.StandardOutPut(podList)
			}

		}
	}
}
