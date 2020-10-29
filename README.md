### apiserver-gin
基于gin的api构键脚手架，项目从实际生产出发，也参考了[golang项目标准布局](https://github.com/golang-standards/project-layout)，还有学习golang过程中见过的比较好的项目结构。
整个项目拉下来，不用做太多改造，就可以实际使用在生产中，特别对于新手学习搭建项目，熟悉项目规划，了解生产项目结构和细节设计感觉会有很大帮助。
本项目依然是按照单体项目布局来做的，后面会完善多模块项目推荐的布局(比如一个项目既有app端server，又有管理后台server)，多项目多模块也是在单体项目上演变的，稍微做下改动而已。

### 下一步计划
- Dockerfile, 消息队列，微服务相关

### 目前整合组件及实现功能
- 加入viper使用yml配置文件来配置项目信息，启动时可根据环境指定不同的配置文件
```html
eg.
go run main.go -config test.yml
```
- 集成gorm 并自定义JsonTime 解决Json序列化和反序列化只支持UTC时间的问题（可以自定义时间格式）  
提供了部分demo，可以按照demo在项目中直接使用。
- 整合redis，开箱即用，直接通过yml文件对redis进行配置
- 整合zap，lumerjack 完善日志输出，日志分割，部分参数可配置。
- 集成jwt，提供demo代码，自定义授权失败成功等的响应格式，跟全局api响应格式统一。
- 实现部分工具类的封装，md5, bcrypt和uuid生成
- 应用统一封装响应格式，基本参照笔者参与的大型项目经验和规范
- 项目全局错误码封装，使go在写业务代码时也能规范统一
- 全局应用常量
- 应用统一入口日志记录中间件实现，类似java中拦截器或者过滤器思想，可以很方便的在入口处记录访问日志。
- 添加makefile，可以使用make 命令进行编译，打包
- 完善了项目版本管理，使用make命令编译后的项目可以方便跟踪线上发布版本
- 其他一些坑，代码中有注释
### 项目使用文档 [查看](https://github.com/xmgtony/apiserver-gin/blob/master/docs/README.md)
