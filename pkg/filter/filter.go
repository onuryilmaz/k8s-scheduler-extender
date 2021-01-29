package filter

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	extenderv1 "k8s.io/kube-scheduler/extender/v1"
)

func Filter(args extenderv1.ExtenderArgs) extenderv1.ExtenderFilterResult {

	filtered := make([]corev1.Node, 0)
	failed := make(extenderv1.FailedNodesMap)

	pod := args.Pod

	for _, node := range args.Nodes.Items {

		side := Flip()
		if side == HEADS {
			filtered = append(filtered, node)
		} else {
			failed[node.Name] = fmt.Sprintf("%s cannot be scheduled to %s: coin is %s", pod.Name, node.Name, side)
		}
	}

	return extenderv1.ExtenderFilterResult{
		Nodes: &corev1.NodeList{
			Items: filtered,
		},
		FailedNodes: failed,
	}

}
