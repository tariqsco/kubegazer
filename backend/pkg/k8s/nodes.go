// Author: Tariq Scott

package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NodeSummary struct {
	Name             string
	IsReady          bool
	KubeletVersion   string
	OSImage          string
	CpuCapacityMilli int64 // e.g., 4000 = 4 cores
	MemCapacityBytes int64 // e.g., 8589934592 bytes
}

// ListNodes lists all nodes in the cluster and returns a slice of NodeSummary
func ListNodes(ctx context.Context, clientset kubernetes.Interface) ([]NodeSummary, error) {
	list, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	summaries := make([]NodeSummary, 0, len(list.Items))
	for _, node := range list.Items {
		summaries = append(summaries, toNodeSummary(node))
	}
	return summaries, nil
}

func toNodeSummary(node corev1.Node) NodeSummary {
	cpu := node.Status.Capacity.Cpu()
	memory := node.Status.Capacity.Memory()

	return NodeSummary{
		Name:             node.Name,
		IsReady:          isNodeReady(node.Status.Conditions),
		KubeletVersion:   node.Status.NodeInfo.KubeletVersion,
		OSImage:          node.Status.NodeInfo.OSImage,
		CpuCapacityMilli: cpu.MilliValue(),
		MemCapacityBytes: memory.Value(),
	}
}

// Helper to extract Ready status from conditions slice
func isNodeReady(conditions []corev1.NodeCondition) bool {
	for _, condition := range conditions {
		if condition.Type == corev1.NodeReady {
			return condition.Status == corev1.ConditionTrue
		}
	}
	return false
}
