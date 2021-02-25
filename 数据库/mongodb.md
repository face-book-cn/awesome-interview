## mongodb数据库

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
	db.文档名.find({"":"","":""})（多个）

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





