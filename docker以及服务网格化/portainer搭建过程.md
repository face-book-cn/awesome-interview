### portainer搭建过程
- 过程分为主节点搭建和从节点搭建
- 两个脚本：分别是主节点脚本文件和从节点脚本文件，直接执行就可以搭建成功

### 搭建流程
```中文
1.首先要保证服务器上安装了docker，下面链接是如何安装docker，如果安装跳过此步骤
2.复制脚本到需要搭建的服务器
3.chmod +x 脚本名字 （执行shell脚本）
4.运行脚本 (./脚本名字.sh)
5.portainer使用教程请移步官网
```

### 主节点 shell脚本：


``` shell
 
    #!/bin/sh
    #auto dhy
    #IP地址
    ipaddress=`hostname -I | awk '{print $1}'`
    
    #开放防火墙端口2375，2377
    firewall-cmd --zone=public --add-port=2375/tcp --permanent   
    firewall-cmd --zone=public --add-port=2377/tcp --permanent
    firewall-cmd --zone=public --add-port=9000/tcp --permanent   
    firewall-cmd --reload         
    
    #设置主机名字
    hostnamectl set-hostname swarm-master-$ipaddress
    
    #拉取swarm
    docker pull swarm
    
    #主节点运行
    docker swarm init --advertise-addr $ipaddress
    
    #启动代理
    docker run -ti -d -p 2375:2375 --restart=always --hostname=$HOSTNAME --name shipyard-proxy -v /var/run/docker.sock:/var/run/docker.sock -e PORT=2375 docker.io/shipyard/docker-proxy:latest
    
    #启动portainer
    docker service create --name portainer -p 9000:9000 --constraint 'node.role == manager' --mount type=bind,src=//var/run/docker.sock,dst=/var/run/docker.sock  portainer/portainer  -H unix:///var/run/docker.sock

```
### 从节点shell脚本

```shell
#!/bin/sh

#IP地址
ipaddress=`hostname -I | awk '{print $1}'`

#开放防火墙端口2375，2377
firewall-cmd --zone=public --add-port=2375/tcp --permanent   
firewall-cmd --zone=public --add-port=2377/tcp --permanent
firewall-cmd --zone=public --add-port=9000/tcp --permanent
firewall-cmd --reload         

#设置主机名字
hostnamectl set-hostname swarm-node-$ipaddress

#拉取swarm
docker pull swarm

#启动代理
docker run -ti -d -p 2375:2375 --restart=always --hostname=$HOSTNAME --name shipyard-proxy -v /var/run/docker.sock:/var/run/docker.sock -e PORT=2375 docker.io/shipyard/docker-proxy:latest
```
