### Go格式化

```go

%v 输出结构体 {10 30}
%+v 输出结构体显示字段名 {one:10 tow:30}
%#v 输出结构体源代码片段 main.Point{one:10, tow:30}
%T 输出值的类型			 main.Point
%t 输出格式化布尔值		 true
%d 输出标准的十进制格式化 100
%b 输出标准的二进制格式化 99 对应 1100011
%c 输出定整数的对应字符  99 对应 c
%x 输出十六进制编码  99 对应 63
%f 输出十进制格式化  99 对应 63
%e 输出科学技科学记数法表示形式  123400000.0 对应 1.234000e+08
%E 输出科学技科学记数法表示形式  123400000.0 对应 1.234000e+08
%s 进行基本的字符串输出   "\"string\""  对应 "string"
%q 源代码中那样带有双引号的输出   "\"string\""  对应 "\"string\""
%p 输出一个指针的值   &jgt 对应 0xc00004a090
% 后面使用数字来控制输出宽度 默认结果使用右对齐并且通过空格来填充空白部分
%2.2f  指定浮点型的输出宽度 1.2 对应  1.20
%*2.2f  指定浮点型的输出宽度对齐，使用 `-` 标志 1.2 对应  *1.20

fmt.Println 打印输出
fmt.Sprintln 把字符返回给一个新的字符串
fmt.Fprintln 把字符返回给一个文件（但不是写入文件）
```

#### go库中常用的方法总结

- 操作数据库 mysql

现在常用xorm框架，如gorm,xorm等等

```go
package main

import (
    //导入数据库用到的库
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
)

func main() {
    //打开数据库连接
	Db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/demo")
	//Db, err := sqlx.Open("mysql", "root:123456@tcp(localhost:3306)/demo")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//延迟关闭数据库连接
	//defer Db.Close()
	//fmt.Printf("mysql ok")
    //执行查询操作
	rows,err := Db.Query("SELECT * FROM t_demo")
	if err != nil{
		fmt.Println("select db failed,err:",err)
		return
	}
	for rows.Next(){
		var name string
		var age string
		err = rows.Scan(&name,&age)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(name,age)
	}
	fmt.Println(result)

	//这种应该是用的很多的，需要结构体
	var user User
	for i := 1; i < 6; i++ {
		err = Db.Get(&user, "SELECT username,age FROM t_demo")
		fmt.Println(user)
	}

	//这种实际就是用一个切片把数据保存起来，不需要结构体
	var userlist []User
	err = Db.Select(&userlist, "select * from t_demo")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userlist)
	for _, data := range userlist {
		fmt.Println(data)
	}

	//Db.Exec()是一个通用的结构。里面可以写删除，更新，增加
	result, err := Db.Exec("update t_demo set username = ? where id = ?", "dhy",2)
	if err != nil{
		fmt.Println(err)
	}
	//增加的id是多少
	//i,_ = result.LastInsertId()
	//影响的行数
	i, _ := result.RowsAffected()
	fmt.Println(i)
}
type User struct {
	Username string
	Age      string
}
```

- 操作 json 数据

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s Student
	str := `{"name":"dhy","age":"26"}`
	i := 1
	fmt.Println(&i)
	fmt.Println(&s)
	//把结构体变为json
	if error := json.Unmarshal([]byte(str), &s)
		error != nil {
		fmt.Println(error)
	}
	fmt.Println(s)
	name := s.Name
	fmt.Println(name)

	s.Name = "56"
	//返回的是一个字节数组
	//把json转为结构体
	bytes, e := json.Marshal(s)
	if e != nil {
		fmt.Println(e)
	}
	i2 := []byte("adad")
	fmt.Print(i2)
	fmt.Printf("%T\n", bytes)
	fmt.Printf("%s\n", bytes)
}

type Student struct {
	Name string
	Age  string
}
```

- 框架 iris

go 拥有许多好用的web框架，iris、beego、gin、echo其用法大体接近

```go
package GETDemo

import (
	"github.com/kataras/iris"
)

func main(){
	application := iris.New()
	application.Get("/dhy/{name}", func(context iris.Context) {
		path := context.Path()
		param := context.URLParam("name")
		application.Logger().Info(param)
		context.WriteString(path)
	})
	application.Run(iris.Addr(":6126"))
}

```

- golang 操作redis

```go
go redis有两套库
	"github.com/garyburd/redigo/redis"
	"github.com/go-redis/redis"

使用第一种
	导入库 go get github.com/garyburd/redigo/redis
---------------------------------------------------------------------
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//连接redis
	conn, e := redis.Dial("tcp", "192.168.38.188:6379")
	checkErr(e)
	reply, e := conn.Do("set", "name", "dhy")
	name, e := redis.String(conn.Do("get", "name"))
	fmt.Println(reply)
	fmt.Println(name)
	//关闭连接
	defer conn.Close()
}
func checkErr(err error) bool {

	if err != nil {
		fmt.Println("错误", err)
		return false
	}
	return true
}


-----------------------------------------------------------------
使用第二种
	导入库 go get github.com/go-redis/redis


package main

import (
	"fmt"
	"github.com/go-redis/redis"
)
func main() {
	//连接redis
	client := redis.NewClient(&redis.Options{
		Addr: "192.168.38.188:6379",
		DB:   0, // use default DB
	})
	s, e := client.Ping().Result()
	checkErr(e)
	fmt.Println(s)

	//redis set 0代表不设置过期时间，这个单位是ns 50000000000ns代表50秒
	i, e := client.Set("dhy", "dhy", 0).Result()
	checkErr(e)
	fmt.Println(i)
    
    //设置过期的key  时间为10秒
	b, e := client.SetNX("hhh", "aaa", 10*time.Second).Result()
	fmt.Println(b)
    //output : true
    
	//redis get
	result, e := client.Get("aa").Result()
	checkErr(e)
	fmt.Println(result)
}
func checkErr(err error) bool {
	if err != nil {
		fmt.Println("错误", err)
		return false
	}
	fmt.Println(err)
	return true
}

```

- go操作消息队列

首先要在服务器上安装nats

```go

//这是最简单的用法，需要可以去nats官网查找其他用法 https://github.com/nats-io/
func main() {
	var url = "nats://192.168.38.209:4222"
	nc, err := nats.Connect(url)
	if err != nil {
		fmt.Println("错误", err)
	}
	//发布消息
	nc.Publish("help", []byte{78, 78, 78})

	//订阅消息
	nc.Subscribe("help", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
	})
	//保证其他go协程可以一直保持连接
	runtime.Goexit()
}
```

