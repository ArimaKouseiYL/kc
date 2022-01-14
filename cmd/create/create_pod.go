package create

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"kc/config"
)

func createPod(name, namespace, image string, port int32) {

	containerName := "ctn-name"

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: namespace},
		Spec: v1.PodSpec{Containers: []v1.Container{
			{Name: containerName,
				Image: image,
				Ports: []v1.ContainerPort{{Name: "port-1", ContainerPort: port}},
			}},
			NodeSelector: map[string]string{"kubernetes.io/hostname": "infra.main.node2"},
		},
	}

	podClient := config.K8sClient.CoreV1().Pods(namespace)

	create, err := podClient.Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		klog.Error(err)
	}
	fmt.Printf("Pod [%s] 创建成功，让我们来查看它的运行状态吧！\n", create.Name)
	fmt.Printf("你可以输入 kc get pod %s -n %s\n", create.Name, create.Namespace)

}
