### kubernetes知识

- 整体介绍

```k8s
ControllerManager负责维护集群的状态，比如故障检测，扩缩容，滚动更新等等。
Scheduler负责资源的调度，按照预定的策略把pod调度到指定的node节点
ETCD 用做已执行存储，pod，service的集群等信息，k8s需要持久化的数据都存储在这个上边。
Kubelet负责维护当前节点上的容器的生命和volumes，网络。
每个Node上可以运行一个kube-proxy，负责service 提供内部的服务发现和负载均衡，为service方法做个落地的功能。
kube-dns负责整个集群的dns服务，这个组件不是必须的，一般通过名字访问比较方便。
dashboard集群数据的GUI界面。
```

- k8s网络

```k8s
CNI
Flannel，Calico，Weave
Pod网络
```

- kube-proxy

```k8s
每个服务，所有的pod给虚拟Ip，虚拟Ip只能在内部访问
ClusterIp

服务暴露到节点，外部的可以通过NodeIp 访问pod
NodePort

kube-DNS
负责集群内部的dns解析，内部之间可以通过名称访问pod。 
```

- k8s

```k8s
master节点（主节点）
    kube-apiserver：提供了HTTP rest接口的关进服务进程，是Kubernetes里所有资源增删改查等操作的唯一入口
    kube-controller-manage：Kubernetes里所有资源对象的自动化控制中心，可以理解为资源对象的“大总管”。
 	kube-scheduler：负责资源的调度（Pod调度）的进程。
    
node节点（从节点）
    kubelet：负责pod对应的同期创建、启动停止等任务，同时与Master节点密切协作，实现集群管理的基本功能。
    kube-proxy：实现Kubernetes Service的通信与负载均衡的重要组件。
    Docker-Engine（docker）：docker引擎，负责本机容器的创建和管理工作。

pod（最基本的部署调度单元）
    每个 Pod 可以由一个或多个业务容器和一个根容器(Pause 容器)组成。一个 Pod 表示某个应用的一个实例。每个pod由一个根容器的pause容器，其他是业务容器。
    k8s为每个pod分配了唯一的IP地址，一个pod里的多个容器共享pod IP
```

- pod状态
```k8s
Pending :
     挂起，这时pod已经被k8s集群接受，但有一个或多个容器镜像尚未创建，等待时间包括调度pod的时间和通过网络下载容器镜像的时间。
Running :
	此时pod已经被绑定到某一个节点上，pod中所有的容器都被创建并且至少有一个容器正在运行或者处于启动或重启状态。 
Succeed :
	此时pod中的所有容器都被成功终止并且不会重启。 
Failed :
	pod中的所有容器都已经被终止，并且至少有一个容器因为失败而终止（容器以非0状态退出）。 
```

- ReplicaSet

>    Pod 副本的抽象，用于解决 Pod 的扩容和伸缩。 

- Deployment

>    Deployment 表示部署，在内部使用ReplicaSet 来实现。可以通过 Deployment 来生成相应的 ReplicaSet 完成 Pod 副本的创建。 

- Service

>    Service 是 Kubernetes 最重要的资源对象。Kubernetes 中的 Service  对象可以对应微服务架构中的微服务。Service 定义了服务的访问入口，服务的调用者通过这个地址访问 Service 后端的 Pod  副本实例。Service 通过 Label Selector 同后端的 Pod 副本建立关系，Deployment 保证后端Pod  副本的数量，也就是保证服务的伸缩性。 



