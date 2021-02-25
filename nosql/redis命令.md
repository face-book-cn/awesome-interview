## redis安装可以使用docker

> docker run --restart=always -p 6379:6379 -d --name redis redis

### 数据结构

- String: 字符串
- Hash: 散列
- List: 列表
- Set: 集合
- Sorted Set: 有序集合
- redis 还可以用来订阅和发布（有消息队列）

### string

string是最基本的数据类型，一个键存储最大512MB

```redis
192.168.38.188:6379> set name dhy
OK
192.168.38.188:6379> get name
"dhy"
```

### hash

Redis hash 是一个键值对集合。
Redis hash 是一个 string 类型的 field 和 value 的映射表，hash 特别适合用于存储对象

```redis
设置person 其中包含name=dhy age=22
hmset person name dhy age 22

获取所有
hgetall person

获取单个
hget person name

获取所有的字段
hkeys person

获取所有值
hvals person

获取给定字段值
hmget person name

添加一个新hash
hset person one 1

删除一个字段
hdel person one
```

### list

Redis列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）

```redis
把pzh cd两个字段放入到address这个list中
lpush address phz cd

遍历address list中的元素 0-5
lrange address 0 5

删除第一个或者最后一个的元素（设置时间）
blpop address cd 1 （左边）
brpop address pzh 1 （右边）

按下标获取元素（从0开始）
lindex address 0

获取元素长度
llen address

移除并获取第一个元素或者最后一个
lpop address
rpop address

通过索引设置元素值
lset address 1 dhy

通过索引移除元素
lrem address 1 hh
```

### set

Redis 的 Set 是 String 类型的无序集合。集合成员是唯一的，这就意味着集合中不能出现重复的数据。
Redis 中集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是 O(1)。

```redis
添加元素
sadd home 1 2 3

获取元素个数
scard home

返回所有元素
smembers home

删除元素
srem home 1 2
```

### sorted set

Redis 有序集合和集合一样也是string类型元素的集合,且不允许重复的成员。
不同的是每个元素都会关联一个double类型的分数。redis正是通过分数来为集合中的成员进行从小到大的排序。
有序集合的成员是唯一的,但分数(score)却可以重复。

```redis
向有序集合添加一个或多个成员，或者更新已存在成员的分数
zadd one 1 2
zadd one 2 redis

遍历元素0-5
zrange one 0 5

获取元素个数
zcard one

返回指定元素下标 0开始
zrank one redis

删除一个或多个元素
zrem one 2 redis
```

### 事务

```redis
取消事务，放弃执行事务块内的所有命令。
DISCARD 

执行所有事务块内的命令。
EXEC 

标记一个事务块的开始。
MULTI 

取消 WATCH 命令对所有 key 的监视。
UNWATCH 

监视一个(或多个) key ，如果在事务执行之前这个(或这些) key 被其他命令所改动，那么事务将被打断。
WATCH key [key ...] 
```

### 其他

```redis
打印字符串
echo dhy

切换数据库
select 1

验证密码是否正确
auth dhy

服务是否运行
ping

查看服务器信息
info

保存数据到硬盘
save
bgsave(后台执行备份)
```

## [redis学习网址](https://www.runoob.com/redis/redis-server.html)