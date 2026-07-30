/*
 * File: backend/pkg/k8s/client.go
 * Author: Tariq Scott
 * Date: July 26, 2026
 * Description: Initializes and manages the Kubernetes client configuration
 *              for KubeGazer. Handles out-of-cluster kubeconfig loading with
 *              fallback to in-cluster service account configuration.
 */

package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// LoadKubeConfig resolves the active kubeconfig, respecting $KUBECONFIG
// and falling back to default loading rules (~/.kube/config, context merging)

func LoadKubeConfig() (*rest.Config, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	overrides := &clientcmd.ConfigOverrides{}

	cfg, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, overrides).ClientConfig()
	if err != nil {
		return nil, err
	}

	// default QPS/Burst is 5/10 - throttles fast once a page fires off
	// pods+nodes+events concurrently. Raising it here, not per-call
	cfg.QPS = 50
	cfg.Burst = 100

	return cfg, nil
}

func NewClientset(cfg *rest.Config) (*kubernetes.Clientset, error) {
	return kubernetes.NewForConfig(cfg)
}
