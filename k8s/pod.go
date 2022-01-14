package k8s

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
)

func StandardOutPut(podList *v1.PodList) {
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

func YamlOutPut(podList *v1.PodList) {
	for _, pod := range podList.Items {

		marshal, err := yaml.Marshal(pod)
		if err != nil {
			return
		}

		fmt.Println(string(marshal))
	}
}
