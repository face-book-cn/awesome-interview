# protobuf

是一种网络传输的二进制数据，跟json、xml、gob一样是用于服务器之间传输的数据格式

- 三个二进制优缺点

`操纵性`：是否需要其他文件来生成（越大越难操作）
`速度`：序列化的速度（越大越快）
`可懂性`：是否利于我们人类阅读（越大越容易读懂）

| 操作性                |         速度          |        可懂性         |
| :-------------------- | :-------------------: | :-------------------: |
| protobuf>xml>gob>json | protobuf>json>gob>xml | json>xml>gob>protobuf |


优点：

```protobuf
- 效率高
- 空间时间的开销小
```

缺点：

```protobuf
- 可读性非常差，就是二进制数据
- 对开发人员不友好
```

## 趋势

在越来越注重高并发、高性能的环境下，越来越多的人用protobuf来作为服务器与服务器之间的传输数据

## 用法

- 安装protobuf（windows、mac、linux版本）
  - 安装好之后控制台输入 `protoc --version`出现版本代表安装ok
- 需要自己写一个proto文件作为序列化文件
- 使用protobuf自带的命令生成文件，如果是java则会生成`.java` go会生成`.pd.go`

语法如下：这只是最基本的东西，如果有需要用到，可以去protobuf官网系统学习一下

```protobuf
//注释可以使用c c++ java风格的注释
syntax = "proto3"; 						//指定版本信息，不指定会使用proto2
package pb;						//后期生成go文件的包名
//message为关键字，作用为定义一种消息类型
message Person{
    string name = 1; //结尾需要加分号
    int32  age = 2 ;
    repeated string emali =3; //数字代表唯一的编号标签值;
    repeated string PhoneNumber =4;
    // repeated为关键字，作用为重复使用 一般在go语言中用切片表示
}

//message为关键字，作用为定义一种消息类型可以被另外的消息类型嵌套使用
message PhoneNumber {
    string number = 1;
    int64 type = 2;
    //enum为关键字，作用为定义一种枚举类型
    enum PhoneType {
        MOBILE = 0;
        ...
    }
}

//每个数据类型都有一个默认值
对于strings，默认是一个空string
对于bytes，默认是一个空的bytes
对于bools，默认是false
对于数值类型，默认是0
//protobuf 对应很多语言的数据类型
如proto的string对应go的string double对应float64
```