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