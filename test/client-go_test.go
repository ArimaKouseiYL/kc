package test

import (
	"context"
	"flag"
	"fmt"
	"github.com/liushuochen/gotable"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	"path/filepath"
	"testing"
)

func TestPod(t *testing.T) {

	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	podList, err := clientSet.CoreV1().Pods("cloudtogo-system").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return
	}

	table, err := gotable.Create("NAME", "STATUS", "NAMESPACE", "NODE")
	if err != nil {
		return
	}

	for _, pod := range podList.Items {
		row := make(map[string]string)
		row["NAME"] = pod.Name
		row["STATUS"] = string(pod.Status.Phase)
		row["NAMESPACE"] = pod.Namespace
		row["NODE"] = pod.Spec.NodeName
		table.AddRow(row)
	}

	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}
	fmt.Println(table)

}

func TestCreatePod(t *testing.T) {

	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "nginx222", Namespace: "example"},
		Spec: v1.PodSpec{Containers: []v1.Container{
			{Name: "nginx",
				Image: "nginx:latest",
				Ports: []v1.ContainerPort{{Name: "port-1", ContainerPort: 80}},
			}},
			NodeSelector: map[string]string{"kubernetes.io/hostname": "infra.main.node2"},
		},
	}

	podClient := clientSet.CoreV1().Pods("example")

	create, err := podClient.Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		klog.Error(err)
	}
	klog.Info(create)

	//clientSet.CoreV1().Pods("example").Create(context.TODO(), pod, metav1.CreateOptions{})

}
