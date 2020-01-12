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

#然后会发现当前目录下多了一个target目录，里面存在编译好的可执行文件，和连配置文件一起打包好的压缩包
直接运行可执行文件
./apidemo-gin

或者指定配置文件运行
./apidemo-gin ./conf/config.yml
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
#### 使用demo代码
1、执行docs目录中的init.sql 对数据库和表初始化  
2、启动项目  
3、请求api地址  
- 用户登录
```text
curl -H "Content-Type: application/json" -X POST  --data '{"username":"xmgtony","password":"123456"}' http://localhost:8080/login
```
响应结果
```text
{
    "request_id": "ce7f7cd542194f4c",
    "err_code": 0,
    "message": "success",
    "data": {
        "expire": "2020-01-13 11:48:46",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Nzg4ODczMjYsImp3dC1rZXkiOnsiaWQiOjEsImNyZWF0ZWQiOiIyMDE5LTEyLTI0IDAwOjQ0OjQxIiwibW9kaWZpZWQiOiIyMDE5LTEyLTI0IDAwOjQ0OjQxIiwibmFtZSI6InhtZ3RvbnkiLCJiaXJ0aGRheSI6IjAwMDEtMDEtMDEgMDA6MDA6MDAifSwib3JpZ19pYXQiOjE1Nzg4MDA5MjZ9.dwBMPXyqBSF-THl7y5ikLOQyEyCVAs-2anEtRAAPB40"
    }
}
```
- 查询用户信息
```text
# 访问时带上一步返回的token信息
curl -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Nzg4ODczMzcsImp3dC1rZXkiOnsiaWQiOjEsImNyZWF0ZWQiOiIyMDE5LTEyLTI0IDAwOjQ0OjQxIiwibW9kaWZpZWQiOiIyMDE5LTEyLTI0IDAwOjQ0OjQxIiwibmFtZSI6InhtZ3RvbnkiLCJiaXJ0aGRheSI6IjAwMDEtMDEtMDEgMDA6MDA6MDAifSwib3JpZ19pYXQiOjE1Nzg4MDA5Mzd9.PmdZyw2NOS0lVG4fNxsUSr2aHbMMhh8Wz7dohGKGStw" http://localhost:8080/v1/user/xmgtony
```
响应结果
```text
{
	"request_id": "3c0fabdd3d2144cd",
	"err_code": 0,
	"message": "success",
	"data": {
		"id": 1,
		"created": "2019-12-24 00:44:41",
		"modified": "2019-12-24 00:44:41",
		"name": "xmgtony",
		"birthday": "0001-01-01 00:00:00"
	}
}
```
- 创建一个新用户
```text
curl -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Nzg4ODczMzcsImp3dC1rZXkiOnsiaWQiOjEsImNyZWF0ZWQiOiIyMDE5LTEyLTI0IDAwOjQ0OjQxIiwibW9kaWZpZWQiOiIyMDE5LTEyLTI0IDAwOjQ0OjQxIiwibmFtZSI6InhtZ3RvbnkiLCJiaXJ0aGRheSI6IjAwMDEtMDEtMDEgMDA6MDA6MDAifSwib3JpZ19pYXQiOjE1Nzg4MDA5Mzd9.PmdZyw2NOS0lVG4fNxsUSr2aHbMMhh8Wz7dohGKGStw" --data '{"name":"xmgtony","password":"123456","birthday" :"2019-08-07 08:09:01"}' http://localhost:8080/v1/user/create
```
响应结果
```text
#失败时
{
	"request_id": "808b1fa6397d4368",
	"err_code": 40006,
	"message": "用户已存在或者提交的信息错误"
}
#成功时
{
	"request_id": "50f60e68551440e3",
	"err_code": 0,
	"message": "success"
}
```
