### conf_model文件

```go
package config

type Config struct {
	Keys  string
	Mongo mongoConf
	Http  httpConf
	Mysql mysqlConf
	Redis redisConf
}

// Mongo
type mongoConf struct {
	Address     string
	MaxPoolSize uint64
}

// Http
type httpConf struct {
	Listen string
}

// Mysql
type mysqlConf struct {
	Address     string
	MaxOpenConn int
	MaxIdleConn int
}

// Redis
type redisConf struct {
	Address     string
	MaxPoolSize int
	Password    string
}

```

### conf.toml文件

```toml
keys="ssssss"

[mongo]
address = "mongodb://dhy:dhy123456@116.85.45.149:27017/admin"
maxPoolSize = 200

[http]
listen = ":6126"

[mysql]
address="root:123456@/aaa?charset=utf8&parseTime=True&loc=Local"
maxIdleConn=10
maxOpenConn=10

[redis]
address = "116.85.45.149:6379"
password = "redis"
maxPoolSize = 200 # 线程池大小
```

### conf.go

```go
package config

import (
	"github.com/spf13/viper"
)

// 传入配置文件类型、所在路径和需要反序列化结构体地址
// 初始化配置只是在启动会初始化。不用担心性能速度问题
func InitConfig(confPath string, conf interface{}) {
	v := viper.New()

	v.SetConfigFile(confPath) // 配置文件的路径

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&conf); err != nil {
		panic(err)
	}

}
```

### conf_test.go

```go
package config

import (
	"fmt"
	"testing"
)

func TestConf(t *testing.T) {
	var c Config
	//var c1 Config
	//var c12 Config

	InitConfig("conf.toml", &c)
	//InitConfig("conf.json", &c1)
	//InitConfig("../init/config.toml", &c12)

	fmt.Println(c)
	//fmt.Println(c1)
	//fmt.Println(c12)
}
```



