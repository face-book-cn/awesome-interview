## 编码规范

### 一、项目目录结构规范

### 文件名命名规范

>目录结构在go中要求不严格，只要方便共同开发的人员能够明白意思即可

```go
├── conf   配置文件目录
├── controller 项目代码，项目逻辑请按照规范命名。
├── main.go 项目主函数
├── models 项目的数据库操作
├── routers 项目路由
├── swagger 自动化Api文档
├── ...
├── ...
```

#### 文件名命名规范

> 驼峰版本：adminUser
>
> 下划线版本：admin_user

### 二、命名规范

#### 包名

> 包名用小写,与外层文件夹名称尽量相同，与标准库不要冲突

#### 接口名

> 接口名以 er 作为后缀，如Reader,Writer

```go
type Reader interface {
    //接口的实现去掉“er”
    Read(p []byte) (n int, err error)
}
```

#### 变量

> 全局变量:采用驼峰命名方式

```go
  var(
      name string
      userAddress int
   )
```

> 局部变量:采用小驼峰命名方式,注意声明局部变量尽量使用 :=

```go
    str := "name"
```

#### 常量

> 全部大写可以使用采用下划线

```
const(
    DIR_NAME= "test"
    STR_NAME = "test"
)
```

#### 函数

> 使用驼峰命名方式，需要包外使用则开头使用大写，否则使用驼峰。

```go
//外部需要访问
func GetList() string()

//外部不需要访问
func getList() string()
```

### 三、注释规范

> go 语言也提供了 /**/ 的块注释和 // 的单行注释两种注释风格， 在我们的项目中为了风格的统一，全部使用单行注释。go 语言自带的 godoc 工具可以根据注释生成文档

#### 包注释

> 每个包都应该有一个包注释，一个位于package子句之前的块注释或行注释。包如果有多个go文件，只需要出现在一个go文件中（一般是和包同名的文件）即可。 包注释应该包含下面基本信息(请严格按照这个顺序，简介，创建人，创建时间）：

```go
// @Title 
// @Description 
// @Author 创建人 创建时间
// @Update  创建人 修改时间
```

#### 结构（接口）注释

> 每个自定义的结构体或者接口都应该有注释说明，该注释对结构进行简要介绍，放在结构体定义的前一行，格式为： 结构体名， 结构体说明。同时结构体内的每个成员变量都要有说明，该说明放在成员变量的后面

```go
// User ， 用户对象，定义了用户的基础信息
type User struct{
    UserName  string `description:"用户名称"`
    Email     string `description:"邮箱"`
}
```

#### 函数 注释

> // @Title  标题
>  // @Description  详细信息
>  // @Auth  创建时间  创建人
>  // @Param      参数类型     参数介绍
>  // @Return  返回类型    "错误信息"

#### 注释风格

> 看个人习惯使用中文或者英文写注释

```go
//准备从 DB 中读取
```

### 代码风格

> go fmt 统一实现风格

### 错误处理

> 1.每个错误都需要处理，不要使用 _ 忽视错误
>
> 2.尽量不要使用panic
>
> 3.错误描述如果是英文必须为小写
>
> 4.直接返回或者日志记录错误