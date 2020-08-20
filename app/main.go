// +----------------------------------------------------------------------
// | Juggle [ 让我们能更好的杂耍 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2020 http://www.XXXXXX.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: dingo <djhui1987@gmail.com>
// +----------------------------------------------------------------------


package main

import (
	"github.com/gin-gonic/gin"
	. "juggle-gin/app/controllers"
	. "juggle-gin/app/middlewares"
	"juggle-gin/src/juggle"
	. "juggle-gin/src/juggle/adapter"
)

func main()  {
	{
		//router := gin.Default()
		//router.Handle("GET","/", func(context *gin.Context) {
		//	context.JSON(200,gin.H{"version":"dev-0.01"})
		//})
		//router.Run(":8080")
	}

	{
		//router := gin.Default()
		//router.Handle("GET","/",NewIndexController().Index())
		//router.Handle("GET","/users",NewUserController().GetUserList())
		//router.Run(":8080")
	}

	{
		//router := gin.Default()
		//NewIndexController(router).Build()
		//NewUserController(router).Build()
		//router.Run(":8080")
	}

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
}
