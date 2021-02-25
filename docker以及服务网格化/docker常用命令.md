## docker常用命令

#### 容器生命周期管理

```docker
run 运行一个容器
start/stop/restart 
kill 杀掉一个容器
rm 删除容器
service 服务操作 常见的ls rm
pause/unpause 暂停容器中所有的进程 以及 恢复容器中所有的进程。
create 创建一个新容器但是不启动
exec 进入容器
```

####  容器操作

```docker
ps （ps -a 查看所有容器）
inspect 获取容器/镜像的元数据。
top 查看容器中运行的进程信息，支持 ps 命令参数。
attach 与exec一样是进入容器的命令，但是退出容器会导致容器停止运行，另外他是同步的，操作同一个容器会阻塞
events 从服务器获取实时事件
logs 日志
wait 阻塞运行直到容器停止，然后打印出它的退出代码。
export 将文件系统作为一个tar归档文件导出到STDOUT。
port 列出指定的容器的端口映射，或者查找将PRIVATE_PORT NAT到面向公众的端口。
cp 复制
```

#### docker自定义容器

```docker
docker commit -a="dhy" -m="del docs tomcat" 容器id dhy/tomcat:1.1
docker commit -a "dhy" -m "del docs tomcat" 容器id dhy/tomcat:1.1 不加 "=" 也是可以的
-a 表明自定义容器的作者
-m 自定义容器的含义
dhy/tomcat 自定义容器的名称
```

#### 本地镜像操作

```docker
images
rmi
tag
build 构建镜像
history 查看指定镜像的创建历史。
//把镜像打成压缩文件，适用迁移 （将指定镜像保存成 tar 归档文件）
save (docker save 镜像名称or容器名称 | gzip > redis-latest.tar.gz)
//把save打好的压缩文件加载成镜像 （导入使用 docker save 命令导出的镜像）
load (docker load -i redis-latest.tar.gz)
import 从归档文件中创建镜像。
```

#### docker信息

````docker
info
version
//docker search redis --filter=stars=50 （–filter=stars=N 类似github starts次数大于多少的）
search 搜索docker仓库里的镜像
login
loginout
````

#### docker数据卷

```docker
docker 数据卷
-v /宿主机文件:/容器文件:ro 
--volumes-from 容器id 表示继承那个容器的数据卷 
数据卷默认是可读可写
:ro 只读不可写
```

#### docker 命令高级用法

```docker
docker stop $(docker ps -q)  停用全部运行中的容器:

docker rm $(docker ps -aq)   删除全部容器：

docker rm `docker ps -a -q`

docker rmi `docker images -q` 删除所有镜像

docker container update --restart=always 容器名字

docker cp docker的复制命令。可以容器复制到宿主机，也可以复制到容器
```

#### docker网络模式

```docker
网络模式：docker的默认网络模式是网桥模式
	host：docker会用宿主机的IP和端口，不会获得一个独立的network  namespace
	container：这个模式指定新创建的容器和已经存在的一个容器共享一个Network Namespace，而不是和宿主机共享
	none：这个模式有自己的独立的网络配置。这种none的也就自己通过exec的方式访问
	bridge：docker会为容器创建一个网络配置，并将docker容器链接到一个虚拟网桥上
```

#### docker高级用法

```docker
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' container_name_or_id可直接获得容器的ip地址如：172.18.0.4

显示所有容器IP地址：docker inspect --format='{{.Name}} - {{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $(docker ps -aq)

docker container update --restart=always 容器名字
```

#### Dockerfile

```dockerfile
FROM 基础镜像，表明镜像基于那个镜像
MAINTAINER 作者
RUN 容器构建运行命令
EXPOSE 对外暴露的端口
WORKDIR 容器进入的默认工作路径
ENV 构建容器中设置环境变量
ADD 文件拷贝加自动解压，自动处理url和解压tar
COPY 将文件和目录复制到容器的文件系统，文件和目录需位于相对于 Dockerfile 的路径中
	COPY src test
	COPY["src","test"]
VOLUME 容器数据卷
CMD 指定容器启动时要运行的命令，多个CMD只有最后一个生效（镜像最后指定会覆盖Dockerfile中的CMD）
ENTRYPOINT 跟CMD一样，但是它会把命令拼接在一起，比CMD命令更强大
```

#### docker-compose

```shell
安装docker-compose
    sudo yum -y install python-pip
    sudo pip install docker-compose

后台运行
	docker-compose up -d
	
查看某个service的compose日志
    #docker-compose logs <service名称>
    docker-compose logs db
    
停止、开始、重启compose服务
	#docker-compose.yml 目录下执行
	docker-compose stop|start|restart
	
kill compose服务
	docker-compose kill
删除compose服务
docker-compose rm

构建服务
	docker-compose build [options] [SERVICE...] 
		–force-rm 删除构建过程中的临时容器。
		–no-cache 构建镜像过程中不使用 cache（这将加长构建过程） 。
		–pull 始终尝试通过 pull 来获取更新版本的镜像。
		
验证 Compose 文件格式是否正确，若正确则显示配置，若格式错误显示错误原因。 
    #校验当前文件夹下的docker-compose.yml
    docker-compose config

此命令将会停止 up 命令所启动的容器，并移除网络 
    #校验当前文件夹下的docker-compose.yml
    docker-compose down

进入指定的容器
    docker-compose exec <service> /bin/sh

获得一个命令的帮助
    docker-compose 命令 help

列出 Compose 文件中包含的镜像
    docker-compose images

暂停、恢复一个服务容器
    docker-compose pause|unpause [SERVICE...]
```

