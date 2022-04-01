### apiserver-gin

基于gin的api服务端脚手架。 gin在Go web开发中是相当受欢迎的，但是gin也是一个轻量级web框架，并不能像其他语言比如java的spring框架具有丰富的生态和标准，在实际开发中需要自己设计和添加一些额外的能力，来完善应用，比如：request_id透传，依赖注入，日志打印分割，session管理，全局错误处理，编译打包等等。

本项目布局为传统的MVC模式，适用于大部分业务api服务端开发。参考了行业流行框架，争取做到每一个文件，每一个目录位置有对应的理论依据支撑。

布局参考[project-layout](https://github.com/golang-standards/project-layout)，该项目非Go官方标准，但是已经是行业主流。


### 理论基础
#### 清洁架构 [(Robert C. Martin)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
![image](https://user-images.githubusercontent.com/8643542/159397149-17f58fba-a3c0-4874-b49a-ae724989af59.png)

按照依赖注入，面向对象编程思想，开闭原则，可拓展，可测性等原则来规划项目。

### 更新日志

*2022-04-01：拆分中间件和路由，不放在一起，便于团队开发时多人维护，比如用户模块的开发人员维护user_router.go，商品模块人员维护goods_router.go互相不影响，便于拓展，详情参考cmd/wire.go*

*2022-03-29：校验器validator支持中文，支持自定义标签，替换了gin默认validator实现，不用每次校验错误后，再翻译成中文，根据配置直接返回中文提示信息，api接口不用处理error翻译。  
具体实现查看pkg/validator包，接口演示查看登录接口。[详细使用文档-点这里](https://github.com/xmgtony/apiserver-gin/blob/master/docs/quick_start.md)*

*2022-03-15：实现日志requestId等透传，使用示例 pkg/log/log_test.go*

*2022-03-11：wire依赖注入工具引入，升级jwt组件。*

*2022-02-22：按照清洁架构及实际项目使用经验重新规划项目结构。*

### 目前整合组件及实现功能

- 加入viper使用yml配置文件来配置项目信息，启动时指定不同的配置文件

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
- 其他一些坑，后续会出一系列配置与使用教程。

#### [使用文档-点这里](https://github.com/xmgtony/apiserver-gin/blob/master/docs/quick_start.md)

### 特别感谢JetBrains对开源项目支持
<a href="https://jb.gg/OpenSourceSupport">
  <img src="https://user-images.githubusercontent.com/8643542/160519107-199319dc-e1cf-4079-94b7-01b6b8d23aa6.png" align="left" height="100" width="100"  alt="JetBrains">
</a>
