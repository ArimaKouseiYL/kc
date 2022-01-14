package delete

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kc/config"
)

func deletePod(name, namespace string) {

	isExist, err := config.K8sClient.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return
	}

	if isExist == nil {
		fmt.Printf("Error from server (NotFound): pods %s not found\n", namespace)
		return
	}

	config.K8sClient.CoreV1().Pods(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	fmt.Printf("Pod [%s] 删除成功！\n", name)

}
