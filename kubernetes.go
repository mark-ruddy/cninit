package cninit

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// getClientRestConfig will use the config of the currently connected kubernetes cluster
// you can confirm what cluster you are connected to using "kubectl config view"
// this config will not be bound to a namespace
func getClientRestConfig() (*rest.Config, error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return config, fmt.Errorf("generating kubernetes rest config: %w", err)
	}
	return config, nil
}

// CreateKubernetesClient will return a *kubernetes.Clientset which can interact with the cluster
func CreateKubernetesClient() (*kubernetes.Clientset, error) {
	config, err := getClientRestConfig()
	if err != nil {
		return nil, fmt.Errorf("while obtaining rest config: %w", err)
	}
	clientset := kubernetes.NewForConfigOrDie(config)
	return clientset, nil
}

// getClientRestConfigNamespace will return a config of the currently connected kubernetes cluster bound to a namespace
func getClientRestConfigNamespace(namespace string) (*rest.Config, error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	namespaceContext := api.Context{Namespace: namespace}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{Context: namespaceContext})
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return config, fmt.Errorf("generating kubernetes namespace rest config: %w", err)
	}
	return config, nil
}

// CreateKubernetesClientNamespace will return a *kubernetes.Clientset which can interact with the cluster
// It will be bound to the namespace specified
func CreateKubernetesClientNamespace(namespace string) (*kubernetes.Clientset, error) {
	config, err := getClientRestConfigNamespace(namespace)
	if err != nil {
		return nil, fmt.Errorf("while obtaining namespace rest config: %w", err)
	}
	clientset := kubernetes.NewForConfigOrDie(config)
	return clientset, nil
}
