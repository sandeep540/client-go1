package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/sandeep/.kube/config", "location of the kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	pods, err := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	ingress, err := clientset.NetworkingV1().Ingresses("").List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are %d ingress in the cluster\n", len(ingress.Items))

	for _, ing := range ingress.Items {
		//Print NAME CLASS HOSTS ADDRESS PORTS AGE
		fmt.Printf("Creation age %v \n", time.Now().Sub(ing.CreationTimestamp.Time).Minutes())
		fmt.Printf("Class Name %s \n", ing.Annotations["kubernetes.io/ingress.class"])
		fmt.Printf(" Name Host IP Port %s %s %d %d \n", ing.Name, ing.Spec.Rules[0].Host, ing.Spec.Rules[0].HTTP.Paths[0].Backend.Service.Port.Number, ing.Spec.Rules[0].HTTP.Paths[0].Backend.Service.Port.Number)
	}

}
