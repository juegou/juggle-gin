// +----------------------------------------------------------------------
// | Juggle [ 让我们能更好的杂耍 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2020 http://www.XXXXXX.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: dingo <djhui1987@gmail.com>
// +----------------------------------------------------------------------

// 核心库
package juggle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

type Juggle struct {
	*gin.Engine
	rg *gin.RouterGroup
	//dba interface{}
	props []interface{}
}

// 初始化
func Ready(engine *gin.Engine) *Juggle {
	jugger := &Juggle{Engine: engine,props:make([]interface{},0)} // 初始化gin引擎
	jugger.Use(ErrorHandler()) // 加载 接收异常 中间件
	return jugger
}

// 加载DB
func (this *Juggle) Beans(beans ...interface{}) *Juggle {
	//this.dba = db
	this.props = append(this.props,beans...)

	return this
}

// 加载中间件
func (this *Juggle) Attach(bag AirBag) *Juggle {
	this.Use(func(ctx *gin.Context) {
		err := bag.Onquest(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(400,gin.H{"msg":err.Error()}) // 中间件报错
		} else {
			ctx.Next()
		}
	})
	return this
}

// 重写Handle
func (this *Juggle) Handle(httpMethod, relativePath string, handlers interface{}) *Juggle {
	// handlers(返回值) 转换成HandlerFunc
	if f := Convert(handlers); f != nil {
		this.rg.Handle(httpMethod,relativePath, f) // 分组路由构建
	}
	return this
}

// 挂载
func (this *Juggle) Mount(group string, controllers ...IController) *Juggle {
	this.rg = this.Group(group) // 定义分组
	for _,controller := range controllers {
		controller.Build(this) // 构建路由

/*
		// 处理控制器加载的DB
		// 通过反射加载控制器想要加载的DB对象
		vclass := reflect.ValueOf(controller).Elem()
		// 数据库结构体是否有字段,假设第一个字段为数据库DB对象
		if vclass.NumField() > 0 {
			if this.dba != nil {
				vclass.Field(0).Set(reflect.New(vclass.Field(0).Type().Elem()))
				vclass.Field(0).Elem().Set(reflect.ValueOf(this.dba).Elem())
			}
		}

		*/


		this.setProps(controller)

	}
	return this
}

// 运行
func (this *Juggle) Go()  {
	config := InitConfig() // 获取系统配置
	this.Run(fmt.Sprintf(":%d",config.Server.Port)) // http服务启动套接字
}


// 控制器中的字段与props中的属性对比
func (this *Juggle) setProps(controller IController)  {
	vclass := reflect.ValueOf(controller).Elem()
	// 遍历控制器结构字段
	for i := 0; i < vclass.NumField(); i++ {
		f := vclass.Field(i) // Field返回结构v的第i个字段。
		// 返回的字段不为nil或者不为指针则不处理
		if !f.IsNil() || f.Kind() != reflect.Ptr {
			continue
		}
		// 属性的类型去props中查找是否有匹配类型
		if p := this.getProps(f.Type()); p != nil {
			f.Set(reflect.New(f.Type().Elem())) // 初始化
			f.Elem().Set(reflect.ValueOf(p).Elem()) // 设置值
		}
	}
}

// 获取props切片中的属性类型 (如：控制器-*adapter.GormAdapter  props-adapter.NewGormAdapter())
// 都是 *GormAdapter类型
func (this *Juggle) getProps(t reflect.Type) interface{} {
	// 遍历props属性
	for _,prop := range this.props {
		if t == reflect.TypeOf(prop) {
			return prop
		}
	}
	return nil
}
