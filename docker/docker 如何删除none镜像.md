删除none的镜像，要先删除镜像中的容器。

要删除镜像中的容器，必须先停止容器。

首先列举所有的镜像：

```shell
docker images
```

如果直接删除带none的镜像，直接报错了，提示先停止容器：

```shell
docker rmi $(docker images | grep "none" | awk '{print $3}')
```



`$ docker stop $(docker ps -a | grep "Exited" | awk '{print $1 }')` //停止容器

`$ docker rm $(docker ps -a | grep "Exited" | awk '{print $1 }')` //删除容器

`$ docker rmi $(docker images | grep "none" | awk '{print $3}')` //删除镜像

```
docker stop $(docker ps -a | grep "Exited" | awk '{print $1 }')
docker rm $(docker ps -a | grep "Exited" | awk '{print $1 }')
docker rmi $(docker images | grep "none" | awk '{print $3}')
```

