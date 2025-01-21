### 设置阿里镜像

### 集成springboot

### 集成springmvc

### linux下安装mysql

### 服务模块拆解

### lombok作用演示

### 集成数据库连接池

### 集成mybatis plus

### 集成druid监控

### 集成mybatis的优化器

### 分工

* common是一个git
* user是一个git

### ape-common

* ape-common-job：分布式任务调度组件
* ape-common-log：日志组件，提供日志切面自动记录及异步日志提升性能
* ape-common-mybatisplus：采用Mybatisplus作为与数据库交互
* ape-common-redis：缓存组件，提供基于redis的操作封装，redis分布式锁，guava的cache工具类
* ape-common-starter：启动类组件，与启动类相关的功能，放到此组件处，目前包含mongoStarter
* ape-common-swagger：swagger组件，提供整体项目访问api的入口及方法文档
* ape-common-test：测试组件，集成springboot-test，及代码单元测试，代码覆盖率，行覆盖率检测
* ape-common-tool：常用的工具类组件，满足业务日常开发的各种需要，保障安全性，低入侵性
* ape-common-web：web组件，提供统一异常处理，web模块转换，统一返回值
* ape-common-websocket：websocket组件，提供一套带鉴权的websocket，引入即用，简单方便
* ape-mail：邮件发送组件

### 项目整体结构优化

01.版本的统一管理
02.common和业务的解耦分割
03.公用配置由业务方决定