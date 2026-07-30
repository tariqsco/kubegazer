/*
 * File: backend/pkg/k8s/pods.go
 * Author: Tariq Scott
 * Date: July 27, 2026
 * Description: Provides pod listing, filtering, and status extraction for
 *              KubeGazer. Maps corev1.Pod resources to lightweight PodSummary
 *              DTOs, handling termination grace periods and container wait states.
 */

package k8s

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodSummary struct {
	Name      string
	Namespace string
	Status    string // TODO -
	Restarts  int32  // TODO -
	Images    []string
	Node      string
	Age       time.Duration
}

// ListPods returns a PodSummary per pod. namespace == "" lists cluster-wide
func ListPods(ctx context.Context, clientset kubernetes.Interface, namespace string) ([]PodSummary, error) {
	list, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	summaries := make([]PodSummary, 0, len(list.Items))
	for _, pod := range list.Items {
		summaries = append(summaries, toPodSummary(pod))
	}
	return summaries, nil
}

func toPodSummary(pod corev1.Pod) PodSummary {
	// Decision #1: Simple status with Terminating check
	status := string(pod.Status.Phase)

	switch {
	case pod.DeletionTimestamp != nil:
		status = "Terminating"

	default:
		status = string(pod.Status.Phase)
	}

	var restarts int32
	var images []string

	// Decision #2: Include init container restarts as well
	for _, containerStatus := range pod.Status.ContainerStatuses {
		restarts += containerStatus.RestartCount
		images = append(images, containerStatus.Image)
	}

	for _, initStatus := range pod.Status.InitContainerStatuses {
		restarts += initStatus.RestartCount
		// images = append(images, initStatus.Image) <-- unsure if needed, may clog UI
	}

	return PodSummary{
		Name:      pod.Name,
		Namespace: pod.Namespace,
		Status:    status,
		Node:      pod.Spec.NodeName,
		Age:       time.Since(pod.CreationTimestamp.Time),
		// Restarts, Images: your loop over pod.Status.ContainerStatuses - decision #2
	}
}
