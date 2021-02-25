<center><h1>docker中redis集群（cluster模式）<h1/><center/>

### docker-compose.yml

```docker
version: '3'
services:
 redis1:
  image: publicisworldwide/redis-cluster
  network_mode: host
  restart: always
  volumes:
   - /data/redis/6180/data:/data
  environment:
   - REDIS_PORT=6180

 redis2:
  image: publicisworldwide/redis-cluster
  network_mode: host
  restart: always
  volumes:
   - /data/redis/6181/data:/data
  environment:
   - REDIS_PORT=6181

 redis3:
  image: publicisworldwide/redis-cluster
  network_mode: host
  restart: always
  volumes:
   - /data/redis/6182/data:/data
  environment:
   - REDIS_PORT=6182

 redis4:
  image: publicisworldwide/redis-cluster
  network_mode: host
  restart: always
  volumes:
   - /data/redis/6183/data:/data
  environment:
   - REDIS_PORT=6183

 redis5:
  image: publicisworldwide/redis-cluster
  network_mode: host
  restart: always
  volumes:
   - /data/redis/6184/data:/data
  environment:
   - REDIS_PORT=6184

 redis6:
  image: publicisworldwide/redis-cluster
  network_mode: host
  restart: always
  volumes:
   - /data/redis/6185/data:/data
  environment:
   - REDIS_PORT=6185
```

### 启动集群命令

```docker
docker run --rm -it inem0o/redis-trib create --replicas 1 192.168.38.131:6180 192.168.38.131:6181 192.168.38.131:6182 192.168.38.131:6183 192.168.38.131:6184 192.168.38.131:6185
```



### 不使用docker-conpose部署

```docker
docker run -d --net=host --name redis_1 -e REDIS_PORT=6180 --restart=always -v /data/redis/6181/data:/data publicisworldwide/redis-cluster
docker run -d --net=host --name redis_2 -e REDIS_PORT=6181 --restart=always -v /data/redis/6182/data:/data publicisworldwide/redis-cluster
docker run -d --net=host --name redis_3 -e REDIS_PORT=6182 --restart=always -v /data/redis/6183/data:/data publicisworldwide/redis-cluster
docker run -d --net=host --name redis_4 -e REDIS_PORT=6183 --restart=always -v /data/redis/6184/data:/data publicisworldwide/redis-cluster
docker run -d --net=host --name redis_5 -e REDIS_PORT=6184 --restart=always -v /data/redis/6185/data:/data publicisworldwide/redis-cluster
docker run -d --net=host --name redis_6 -e REDIS_PORT=6185 --restart=always -v /data/redis/6186/data:/data publicisworldwide/redis-cluster

#会自动确定
echo yes | docker run --rm -i inem0o/redis-trib create --replicas 1 172.17.0.1:6180 172.17.0.1:6181 172.17.0.1:6182 172.17.0.1:6183 172.17.0.1:6184 172.17.0.1:6185
```

