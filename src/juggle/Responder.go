// +----------------------------------------------------------------------
// | Juggle [ 让我们能更好的杂耍 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2020 http://www.XXXXXX.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: dingo <djhui1987@gmail.com>
// +----------------------------------------------------------------------

// 处理响应数据
package juggle

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

func init()  {
	ResponderList = []Responder{
		new(StringResponder),
		new(ModelResponder),
		new(ModelsResponder),
	}
}

var ResponderList []Responder

type Responder interface {
	ResponderTo() gin.HandlerFunc
}

// 数据转换成gin.HandlerFunc,启动时即执行了，所以不会影响性能
// 思路：通过反射获取返回值的类型
//      与ResponderList中的处理类型相比较
//      相同类型则返回
func Convert(handler interface{}) gin.HandlerFunc {
	handlerReflect := reflect.ValueOf(handler)
	for _,resp := range ResponderList {
		respValue := reflect.ValueOf(resp).Elem() // interface value 包含的值
		// 判断接口的类型是否可以转换
		if handlerReflect.Type().ConvertibleTo(respValue.Type()) {
			respValue.Set(handlerReflect) // 可以转换，将handlerReflect赋值给respValue
			return respValue.Interface().(Responder).ResponderTo() // 等同于 var i interface {} =（v的基础值）
		}
	}
	return nil
}


type StringResponder func(*gin.Context) string

func (this StringResponder) ResponderTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(200,this(ctx))
	}
}

type ModelResponder func(*gin.Context) IModel

func (this ModelResponder) ResponderTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{"info":this(ctx)})
	}
}


type ModelsResponder func(*gin.Context) Models

func (this ModelsResponder) ResponderTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type","application/json")
		ctx.Writer.WriteString(string(this(ctx)))
	}
}