### 使用脚本安装docker以及docker-compose
```shell
#!/bin/sh

#移除旧版本docker
yum remove docker \
    docker-client \
    docker-client-latest \
    docker-common \
    docker-latest \
    docker-latest-logrotate \
    docker-logrotate \
    docker-selinux \
    docker-engine-selinux \
    docker-engine

#安装一些必要的系统工具
yum install -y yum-utils device-mapper-persistent-data lvm2

#添加软件源信息
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

#更新 yum 缓存
yum makecache fast

#安装 Docker-ce
yum -y install docker-ce

#启动 Docker 后台服务
systemctl start docker

#docker加入开机自启动
systemctl enable docker

echo '"registry-mirrors": ["https://gz1chk08.mirror.aliyuncs.com"] }' > /etc/docker/daemon.json

systemctl daemon-reload
systemctl restart docker

#下载docket-compose
curl -L https://github.com/docker/compose/releases/download/1.17.1/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose

#修改权限
chmod +x /usr/local/bin/docker-compose

```



### 也可以使用官方脚本安装

```shell
curl -fsSL get.docker.com -o get-docker.sh
sh get-docker.sh --mirror Aliyun
```

