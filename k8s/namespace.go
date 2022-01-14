package k8s

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kc/config"
)

func GetNamespaceNames() []v1.Namespace {

	list, err := config.K8sClient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil
	}
	return list.Items
}
