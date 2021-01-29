package prioritize

import (
	extenderv1 "k8s.io/kube-scheduler/extender/v1"
)

func Prioritize(args extenderv1.ExtenderArgs) extenderv1.HostPriorityList {

	hostPriority := make(extenderv1.HostPriorityList, 0)

	for _, node := range args.Nodes.Items {
		hostPriority = append(hostPriority, extenderv1.HostPriority{
			Host:  node.Name,
			Score: Roll(),
		})
	}

	return hostPriority

}
