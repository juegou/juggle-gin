// +----------------------------------------------------------------------
// | Juggle [ 让我们能更好的杂耍 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2020 http://www.XXXXXX.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: dingo <djhui1987@gmail.com>
// +----------------------------------------------------------------------

// 统一错误处理
package juggle

import "github.com/gin-gonic/gin"

// 接收异常 中间件
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				ctx.AbortWithStatusJSON(400,e)
			}
		}()

		ctx.Next()
	}
}

// 错误统一处理
func Error(err error, msg ...string)  {
	if err == nil {
		return
	} else {
		errMsg := err.Error()
		if len(msg) > 0 {
			errMsg = msg[0]
		}
		panic(errMsg)
	}
}
