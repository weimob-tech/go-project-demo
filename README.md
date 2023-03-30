# go-project-demo
微盟云开放平台的GO项目Demo工程，开发者可以参考该Demo工程实现业务开发，快速切入。整个Demo工程符合微盟云开发平台的GO工程的标准规范，支持微盟云开发平台的构建、发布、监控、治理...

## 介绍
### 功能列表
* GO项目工程结构
* 项目配置管理
	* 项目属性配置
	* 请求路由管理
	* 扩展点注册表
	* 消息订阅表

### 项目结构

```
│  .gitignore
│  go.mod   #项目依赖
│  go.sum
│  LICENSE.md
│  README.md
├─api
├─cmd # 构建入口，WDP 构建默认构建 cmd/ 目录
│  └─main.go
├─configs     # 项目配置文件目录
│ └─application.properties
├─internal    # 内部业务包，不对外暴露
│  ├─app       # 内部全局功能目录
│  │  ├─app.go # 注册 SPI、消息
│  │  └─route #全局注册的 route
│  │    └─route.go
│  └─pkg      # 可以再根据业务细化，但是注意不要循环依赖，否则放入app包中
│    ├─ability # 消息、扩展点实现
│    │  ├─msg
│    │  └─spi
│    └─controller
├─scripts 
└─test   # 集成测试目录 
```


## 快速开始
1. 下载go-project-demo的代码，可以有以下几种方式
	*  在微盟云开发平台创建的GO容器应用，默认会生成工程结构
	*  git clone git@github.com:weimob-tech/go-project-demo.git
	*  也可以在github下载demo的zip包
2. 修改项目名与go.mod里的module
3. 添加项目的go module的依赖包，默认的依赖包有
	* [weimob-tech/go-project-boot](https://github.com/weimob-tech/go-project-boot)，[使用文档](https://github.com/weimob-tech/go-project-boot/blob/master/README.md)
	* [weimob-tech/go-ability-sdk](https://github.com/weimob-tech/go-ability-sdk) ，[使用文档](https://github.com/weimob-tech/go-ability-sdk/blob/master/README.md)
4. 安装包，使用go命令
	* `GO111MODULE=on GOPROXY=https://goproxy.cn,direct go mod tidy`
5. 本地调试启

      ``` bash
       #linux
       go run ./cmd

       #window
       go run .\cmd
      ```  
6. 生产环境运行，可以在微盟云开发平台进行构建镜像，并发布到容器集群

## 使用文档
* [能力文档](http://doc.weimobcloud.com/list?tag=2396&menuId=19&childMenuId=1&isold=2)
* [开发者入驻](http://doc.weimobcloud.com/word?menuId=46&childMenuId=47&tag=2970&isold=2)
* [应用开发](http://doc.weimobcloud.com/word?menuId=53&childMenuId=54&tag=2488&isold=2)
* [GO工程](http://doc.weimobcloud.com/word?menuId=53&childMenuId=54&tag=2488&isold=2)

## 贡献方法
* 申请加入weimob_tech

## 联系我们
* Weimob-tech@weimob.com
