### apiserver-gin

基于gin的api服务端脚手架。 对于新手学习使用gin搭建go服务端项目，熟悉项目规划，了解生产项目结构和细节设计感觉会有很大帮助。 本项目依然是按照单体项目布局来做的，后面会完善多模块项目推荐的布局(
比如一个项目既有app端server，又有管理后台server)。

### 理论基础
#### 清洁架构 [(Robert C. Martin)](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
![image](https://user-images.githubusercontent.com/8643542/159397149-17f58fba-a3c0-4874-b49a-ae724989af59.png)

按照依赖注入，面向对象编程思想，开闭原则，可拓展，可测性等原则来规划项目。

### 更新日志

***2022-03-29：校验器validator支持中文，支持自定义标签，替换了gin默认validator实现，不用每次校验错误后，再翻译成中文，根据配置直接返回中文提示信息，api接口不用处理error翻译。  
具体实现查看pkg/validator包，接口演示查看登录接口。[详细使用文档-点这里](https://github.com/xmgtony/apiserver-gin/blob/master/docs/quick_start.md)***

***2022-03-15：实现日志requestId等透传，使用示例 pkg/log/log_test.go***

***2022-03-11：wire依赖注入工具引入，升级jwt组件。***

***2022-02-22：按照清洁架构及实际项目使用经验重新规划项目结构。***

### 目前整合组件及实现功能

- 加入viper使用yml配置文件来配置项目信息，启动时指定不同的配置文件

```html
*** 直接运行 ***
eg. 在cmd目录下执行（也可以在其他目录执行，注意配置文件路径，默认寻找当前执行路径下conf目录中的config.yml文件）
go run . -c ../conf/config.yml

*** 使用make ***
//1、打包（Linux/MacOS 下），在项目目录下执行make命令，打好的包在target目录下
make 
//2、删除已打的包
make clean
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