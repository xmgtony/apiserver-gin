### 项目结构
```text
├── conf            # 配置文件目录
├── docs            # 文档
├── handler         # 定义handler的目录
│   ├── check       # 服务器启动自检的handler所在目录
│   └── user        # 用户相关的handler(一个简单的用户操作demo)
│       └── v1      # 用户handler的版本目录
├── model           # model目录，一般和数据库表结构对应
│   └── user        
├── pkg             # 项目核心代码目录
│   ├── cache       # 缓存相关，目前只实现了redis
│   ├── config      # 跟配置相关的代码包，比如解析yml配置文件
│   ├── constant    # 常量目录
│   ├── errcode     # 自定义的错误处理相关代码目录
│   ├── log         # 日志相关
│   ├── response    # 跟http请求响应相关代码目录
│   ├── time        # 时间处理
│   └── version     # 版本
├── router          # 路由
│   └── middleware  # 中间件
├── scripts         # 项目脚本，比如启动脚本
├── service         # 服务层代码目录
│   └── user        
│       └── v1
└── tools           # 工具包
    └── security
```
### 开始使用
#### 项目依赖
```text
mysql(可选)，redis(可选)，go(1.12及以上版本)
```
最简单的使用方式：
```text
# 1、clone 源码

git clone https://github.com/xmgtony/apidemo-gin

# 2、打开conf目录里面的配置文件，修改mysql和redis配置

# 3、在linux下执行make命令

make

#然后会发现当前目录下多了一个target目录，里面存在编译好的可执行文件，和连配置文件一起打包好的压缩包。
```
也可以直接运行main.go, 或者使用go build main.go 编译运行。

如果不需要mysql和redis，请在main.go中注释掉相关的加载，这样可以不需要mysql和redis直接启动。
```go
// 加载配置文件
config.Load(configFile)
// 初始化Redis Client
cache.RedisInit()
defer cache.RedisClose()
// 初始化数据库信息
model.DBInit()
defer model.DBClose()
// 初始化logger
log.LoggerInit()
```
注释后
```go
// 加载配置文件
config.Load(configFile)
// 初始化Redis Client
//cache.RedisInit()
//defer cache.RedisClose()
// 初始化数据库信息
//model.DBInit()
//defer model.DBClose()
// 初始化logger
log.LoggerInit()
```
