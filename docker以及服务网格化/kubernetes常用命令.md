### kubernetes常用命令

	查看节点  kubectl get nodes
	查看pod   kubectl get pods -n kube-system   |   kubectl get pods --namespace=kube-system | kubectl get pods --all-namespace
	查看pod中的images  kubectl get pod
	查看日志  kubectl logs -n kube-system [podname]
	查看详细的pod信息	kubectl get pods -A -o wide
	查看服务	kubectl get svc -n kube-system
	删除节点  kubectl delete node nodename
	查看token   kubeadm token list
	建立一个主节点 kubeadm init
	加入主节点  kubeadm join
	还原所有的kubeadm join 和 kubeadm init 操作    kubeadm reset
	重启kubernetes   systemctl restart kubelet.service

