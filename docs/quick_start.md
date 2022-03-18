#### 克隆代码到本地

```shell
git clone https://github.com/xmgtony/apiserver-gin.git
```

#### 创建数据库和表

sql文件在当前项目docs/init.sql目录下。 因为在实际项目开发中，都是先确定好数据模型，设计好表才开始开发代码，并且要提供原始的sql脚本给code reviewer或者DBA人员审核。
所以这里并没有使用gorm中提供的自动根据struct创建表功能来创建表。另外如果切换gorm为sqlx或者其他数据库访问库，提供原始sql也是很有必要的。

#### 配置项目的配置文件

项目的配置文件在conf目录下，提供了一个模板文件config.yml, 你可以直接修改该文件或者复制一份。第一次启动只需要配置数据库链接信息。把以下信息改为你的数据库信息。

```yaml
# 数据库配置（mysql）
database:
  dbname: apiserver_gin #数据库名称
  host: 127.0.0.1 #数据库服务地址
  port: 3306 #数据库端口
  username: root #用户名(实际线上不可使用root,请授权一个账户)
  password: 123456 #密码
```

#### 启动服务

项目的入口函数在cmd的main.go中。所以我们只需要在cmd目录下执行以下命令。  
（注意：当前cmd下有多个go文件，包括main.go 都是在package main下，所以不能使用go run main.go方式运行，因为不会加载同包以下其他go文件，会报错。）

```shell
>$ go run . -c ../conf/config.yml
```

不出意外控制台会输出以下信息：

```shell
2022/03/17 23:53:12 The application configuration file is loaded successfully!
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> apiserver-gin/internal/handler/ping.Ping.func1 (8 handlers)
[GIN-debug] POST   /login                    --> apiserver-gin/internal/handler/v1/user.(*UserHandler).Login.func1 (8 handlers)
[GIN-debug] GET    /v1/user                  --> apiserver-gin/internal/handler/v1/user.(*UserHandler).GetUserInfo.func1 (9 handlers)
[GIN-debug] POST   /v1/user/login            --> apiserver-gin/internal/handler/v1/user.(*UserHandler).Login.func1 (9 handlers)
2022-03-17 23:53:12.953	INFO	middleware/logger.go:26	New request start	{"appName": "apiserver-gin", "request_id": "0c149ba41f974f47", "host": "127.0.0.1", "host": "127.0.0.1", "path": "/ping", "method": "GET", "body": ""}
2022-03-17 23:53:12.953	INFO	middleware/logger.go:37	New request end	{"appName": "apiserver-gin", "request_id": "0c149ba41f974f47", "host": "127.0.0.1", "path": "/ping", "cost": "131.344µs"}
[GIN] 2022/03/17 - 23:53:12 | 200 |     357.782µs |       127.0.0.1 | GET      "/ping"
2022-03-17 23:53:12.953	INFO	server/http.go:81	server started success! port: :8080	{"appName": "apiserver-gin"}
```

#### 使用make打包

在linux/MacOS下可以使用make来打包项目

```shell
# 在项目目录下执行make命令，会自动在当前目录寻找Makefile文件
>$ make
# 控制台会输出
cp -r conf  target/
cp ./scripts/startup.sh target/
cd target/ && tar -zcvf apiserver-gin.tar.gz *
a apiserver-gin
a conf
a conf/config.yml
a startup.sh
```

然后我们可以去target目录查看生成的文件信息

```shell
>$  tree target
target
├── apiserver-gin
├── apiserver-gin.tar.gz
├── conf
│   └── config.yml
└── startup.sh
```

其中有可执行二进制文件 apiserver-gin，配置文件 config.yml 和启动脚本startup.sh以及把这三者打包一起的zip成果包apiserver-gin.tar.gz

我们可以直接使用启动脚本启动项目

```shell
>$  ./startup.sh apiserver-gin
# 控制台输出
已启动apiserver-gin...
进程pid：2477
```

我们可以查看下进程是否启动

```shell
>$ ps aux | grep apiserver-gin
```

启动后可以访问 http://127.0.0.1:8080/v1/user

```shell
>$ curl http://127.0.0.1:8080/v1/user

{"request_id":"1c928f4b538147f1","err_code":40001,"message":"invalid token, token is empty"}
```

返回的信息告诉我们需要一个token，我们需要先登录拿到token, 初始化用户名和密码分别为：xmgtony，123456

```shell
>$ curl -H "Content-Type:application/json" -X POST -d '{"username":"xmgtony","password":"123456"}' http://127.0.0.1:8080/login

{
	"request_id": "6e513fddbf694d84",
	"err_code": 0,
	"message": "success",
	"data": {
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJpc3MiOiJhcGlzZXJ2ZXItZ2luIiwiZXhwIjoxNjQ4MjE4ODQ5LCJpYXQiOjE2NDc2MTQwNDl9.0dCx7ciHipYYUWlTmGxvUQpTp0vf79XRp5kQWQJTz04",
		"expire_at": "2022-03-25 22:34:09"
	}
}
```

可以看到成功后返回了token以及token的过期时间，然后把token放在请求的header中，再次查询用户信息

```shell
curl -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJpc3MiOiJhcGlzZXJ2ZXItZ2luIiwiZXhwIjoxNjQ4MjE4ODQ5LCJpYXQiOjE2NDc2MTQwNDl9.0dCx7ciHipYYUWlTmGxvUQpTp0vf79XRp5kQWQJTz04" http://127.0.0.1:8080/v1/user 

{"request_id":"3522cfa70f234fce","err_code":0,"message":"success","data":{"id":1,"created":"2022-03-06 12:45:32","modified":"2022-03-06 12:45:32","name":"xmgtony","birthday":"0001-01-01 00:00:00"}}
```

清除打包信息，会删除target目录

```shell
>$ make clean 
```

#### 后续会推出相关教程，介绍设计理念，以及在实际企业中的开发规范。