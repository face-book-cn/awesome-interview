## mongo数据库

- 优势
  - 文档类型数据库，无需想关系型数据库一样建表。直接使用实体类即可自动建表
  - 百万级以上数据比关系型数据库拥有更加快速的查询速度

- 简单使用

> 所有都可以用 db.getCollection("service_monitoring").find()

```mongodb
#查询所有 同select * from t_demo
	db.文档名.find()
	
#查询条数 select count(1) from t_demo
	db.文档名.find().count()
	
#查询指定 select * from t_demo where id = 1
	db.文档名.find({"id":"1"})（单个）
	db.文档名.find({"":"","":""})（多个，相当于and操作）

#插入 insert into t_demo(name) value(111)
	db.文档名.insert({"id":"2"})
	
#删除 delete from t_demo where id = xxx
	db.文档名.remove();#删除所有数据
    db.文档名.remove({"sex":"女"});#按照条件删除
    db.文档名.remove({"name":"dhy"},2);#删除几条
    
#更新	update t_demo set name = dhy where name = 111
	db.文档名.update({"id":"2"},{"name":"5"})
```

- 排序 分页

> 1是升序，-1是降序

```mongo
#排序 （按照updateTime降序排序）
	db.service_monitoring.find().sort({"updateTime":-1})
	
#排序
	db.文档名.find().sort({"age":-1})（按照sort里面key的值排序，1为正序，-1为倒序）
	
#查询分页 select * from t_demo limit 10
	db.文档名.find().limit(10) || db.文档名.find().limit(10).skip(10)（skip代表跳过多少）
```

- 比较

| 等于       | `{<key>:<value>`}        | `db.col.find({"by":"one"}).pretty()`        | `where by = 'one'`  |
| ---------- | ------------------------ | ------------------------------------------- | ------------------- |
| 小于       | `{<key>:{$lt:<value>}}`  | `db.col.find({"likes":{$lt:50}}).pretty()`  | `where likes < 50`  |
| 小于或等于 | `{<key>:{$lte:<value>}}` | `db.col.find({"likes":{$lte:50}}).pretty()` | `where likes <= 50` |
| 大于       | `{<key>:{$gt:<value>}}`  | `db.col.find({"likes":{$gt:50}}).pretty()`  | `where likes > 50`  |
| 大于或等于 | `{<key>:{$gte:<value>}}` | `db.col.find({"likes":{$gte:50}}).pretty()` | `where likes >= 50` |
| 不等于     | `{<key>:{$ne:<value>}}`  | `db.col.find({"likes":{$ne:50}}).pretty()`  | `where likes != 50` |

- or 和 and

```mongodb
db.col.find({$or:[{"by":"one"},{"title": "MongoDB one"}]}) or
相当于where by = or title = 

db.col.find({"likes": {$gt:50}, $or: [{"by": "one"},{"title": "MongoDB one"}]}) and和or联合使用
相当于where likes>50 and (by = or title = ) 
```

- 模糊查询

```
查询 title 包含"教"字的文档：
db.col.find({title:/教/})

查询 title 字段以"教"字开头的文档：
db.col.find({title:/^教/})

查询 titl e字段以"教"字结尾的文档：
db.col.find({title:/教$/})
```

