package delete

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kc/config"
)

func deletePod(name, namespace string) {
	err := config.K8sClient.CoreV1().Pods(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		fmt.Printf("Error from server (NotFound): pods %s not found\n", namespace)
		return
	}
	fmt.Printf("Pod [%s] 删除成功！\n", name)

}
