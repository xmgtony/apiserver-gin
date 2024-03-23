### apiserver-gin

基于gin的api服务端脚手架。 gin在Go web开发中是相当受欢迎的，但是gin也是一个轻量级web框架，并不能像其他语言比如java的spring框架具有丰富的生态和标准，在实际开发中需要自己设计和添加一些额外的能力，来完善应用，比如：request_id透传，依赖注入，日志打印分割，session管理，全局错误处理，编译打包等等。

本项目布局为传统的MVC模式，适用于大部分业务api服务端开发。参考了行业流行框架，争取做到每一个文件，每一个目录位置有对应的理论依据支撑。

布局参考[project-layout](https://github.com/golang-standards/project-layout)，该项目非Go官方标准，但是已经是行业主流。

### 环境配置
1. 安装go sdk, 建议使用go1.21版本以上，[Go官网下载地址](https://go.dev/dl/)
2. 安装好go sdk以后，需要安装Google wire，依赖注入工具。用于生成依赖注入代码
    ```shell
    $ go install github.com/google/wire/cmd/wire@latest
    ```
3. 设置go的模块代理为国内镜像地址，避免拉取依赖失败
    ```shell
    $ go env -w GO111MODULE=on
    $ go env -w GOPROXY=https://goproxy.cn,direct
    ```

### 运行项目
1. 克隆项目后首先拉取依赖
    ```shell
    $ go mod tidy
    ```
2. 命令行运行
    ```shell
    # 直接运行
    # eg. 在cmd目录下执行（也可以在其他目录执行，注意配置文件路径，默认寻找当前执行路径下conf目录中的config.yml文件）
    >$ go run . -c ../conf/config.yml
    
    # 使用make
    # 1、打包（Linux/MacOS 下），在项目目录下执行make命令，打好的包在target目录下
    >$ make 
    # 2、删除已打的包
    >$ make clean
    ```
   开发工具goland和vscode中运行请自行查找资料，非常简单，不要忘了指定配置文件目录不然找不到配置文件。

### [详细使用文档-点这里](https://github.com/xmgtony/apiserver-gin/blob/master/docs/quick_start.md)

### 理论基础

#### 清洁架构 [(Robert C. Martin)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

![image](https://user-images.githubusercontent.com/8643542/159397149-17f58fba-a3c0-4874-b49a-ae724989af59.png)

按照依赖注入，面向对象编程思想，开闭原则，可拓展，可测性等原则来规划项目。

### 目前整合组件及实现功能

- 加入viper使用yml配置文件来配置项目信息，启动时指定不同的配置文件
- 优雅停机实现，停机时清理资源。
- 集成gorm 并自定义JsonTime 解决Json序列化和反序列化只支持UTC时间的问题（可以自定义时间格式）  
  提供了部分demo，可以按照demo在项目中直接使用。
- 整合redis，开箱即用，通过yml文件对redis进行配置
- 整合zap，lumerjack 完善日志输出，日志分割。
- 集成jwt，提供demo代码，自定义授权失败成功等的响应格式，跟全局api响应格式统一
- 实现session管理
- md5, bcrypt和uuid生成工具包
- 应用统一封装响应格式，参照笔者参与的大型项目经验和规范。
- 项目全局错误码封装，go的error封装。
- 应用统一入口日志记录中间件实现，日志log_id透传。
- 添加makefile，可以使用make 命令进行编译，打包。
- 完善了项目版本管理，使用make命令编译后的项目可以方便跟踪线上发布版本
- 更多功能会根据个人实际开发中的经验添加，不过度封装，保持简单。

### 特别感谢JetBrains对开源项目支持
<a href="https://jb.gg/OpenSourceSupport">
  <img src="https://user-images.githubusercontent.com/8643542/160519107-199319dc-e1cf-4079-94b7-01b6b8d23aa6.png" align="left" height="100" width="100"  alt="JetBrains">
</a>
