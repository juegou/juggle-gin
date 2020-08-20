# Changelog （开发日志）

此项目的所有开发日志都记录在此文件中

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.0.1] - 2020-08-13
### Added
- 创建项目
- 使用gin框架构建一个简单的API
```go
	{
		router := gin.Default()
		router.Handle("GET","/", func(context *gin.Context) {
            // 如果业务多了，乱码凌乱，我们需要封装一个控制器来调度
			context.JSON(200,gin.H{"version":"dev-0.01"})
		})
		router.Run(":8080")
	}
```

- 代码多了，main.go入口文件代码太多不易维护
- 将业务函数gin.HandlerFunc封装到controllers下的控制器文件中，这样让入口函数更清洁
```go
	{
		router := gin.Default()
		router.Handle("GET","/",NewIndexController().Index())
		router.Handle("GET","/users",NewUserController().GetUserList())
		router.Run(":8080")
	}
```

- 将gin的构建路由函数Handle也封装到controllers下的控制器文件中，
  这样只需要NewIndexController(gin.New()).Build()即完成了路由构建
```go
	{
		router := gin.Default()
		NewIndexController(router).Build()
		NewUserController(router).Build()
		// 如果控制器多了，代码会冗余
		// 各个控制器之间没有约束(没有接口规范)
		// 再进化一下
		router.Run(":8080")
	}
```
  这样业务函数与路由构建都封装在了控制器中

## [v0.0.2] - 2020-08-14
### Added  
- 创建核心包juggle，将http服务封装进来
- 所有控制器继承IController
  这样 main.go只需这样调用即可
```go
	{
		juggle.Ready(gin.New()).Mount(
			NewUserController(),
			).Go()
	}
```

## [v0.0.3]
### Added
- 实现gin的路由分组，gin的实现方式
```go
	router := gin.Default()
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}
	router.Run(":8080")
```

main.go只需这样调用即可
```go
	{
		juggle.Ready(gin.New()).
			Mount("v1",NewUserController()).
			Mount("v2",NewUserController()).
			Go()
	}
```


## [v0.0.4]
### Added
- 实现gin的中间件，gin的实现方式
```go
	r := gin.New()
	r.Use(gin.Logger())
```
main.go只需这样调用即可
```go
	{
		juggle.Ready(gin.New()).
			Attach(NewUserMiddleware()).
			Mount("v1",NewUserController()).
			Mount("v2",NewUserController()).
			Go()
	}
```


## [v0.0.5]
### Added
- 控制器业务代码返回的都是gin.HandlerFunc
- 让控制器业务代码支持返回string
- 让控制器支持返回实体类
- 让控制器支持返回实体类切片

  Responder: 控制器返回值处理,统一转换成gin.HanderFunc


## [v0.0.6]
### Added
- Error的统一处理
- 整合ORM (Gorm、Xorm等)
- 简单的bean注入

main.go调用
```go
	{
		juggle.Ready(gin.New()).
			Beans(NewGormAdapter(),NewXormAdapter()).
			Attach(NewUserMiddleware()).
			Mount("v1",
				NewIndexController(),
				NewNewsController(),
				NewUsersController()).
			Go()
	}
```


## [v0.0.7]
### Added
- 集成yaml的配置、加载server配置
- 使用注解方式读取配置
- Bean Factory实现注入(inject)
  把配置注册到bean工厂中



## [v0.0.8]
### Added
- 模板渲染 


## [v0.0.9]
### Added
- 添加协程任务
- 添加定时任务


## [v0.1.0]
### Added
- 更新目录结构

```
├── app                          // 项目文件夹
│   ├── config                   // 配置文件
│   ├── controllers              // 控制器代码
│   ├── middlewares              // 中间件代码
│   ├── models                   // 数据库模型代码
│   ├── application.yaml         // 系统配置文件
│   └── main.go                  // 入口文件
├── src                          // 源代码
│   └── juggle                   // 核心库
│       ├── adapter              // 数据库连接适配器
│           ├── GormAdapter      // Gorm
│           └── XormAdapter      // Xorm
│       ├── Airbag.go            // 规范中间件
│       ├── Error.go             // 错误统一处理
│       ├── Helper.go            // 函数库帮助文件
│       ├── IController.go       // 规范控制器
│       ├── IModel.go            // 规范Model
│       ├── Juggle.go            // 核心库、Http服务包装
│       ├── Responder.go         // 返回值处理
│       └── SysConfig            // 配置文件解析
├── .gitignore                   // git 忽略项
└── go.mod                       // go.mod
```