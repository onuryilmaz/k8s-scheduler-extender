$ minikube mount $(pwd):/etc/k8s-scheduler-extender

$ minikube ssh
§ sudo su

$ cd /etc/k8s-scheduler-extender/manifests
$ cp kube-scheduler-extender.yaml /etc/kubernetes/manifests/kube-scheduler-extender.yaml
$ cp kube-scheduler-config.yaml /etc/kubernetes/kube-scheduler-config.yaml

$ cp kube-scheduler.yaml /etc/kubernetes/manifests/kube-scheduler.yaml

$ systemctl restart kubelet