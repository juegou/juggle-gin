package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"juggle-gin/app/models"
	"juggle-gin/src/juggle"
	"juggle-gin/src/juggle/adapter"
)

type UsersController struct {
	*adapter.GormAdapter
}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (this *UsersController) GetUserDetail(ctx *gin.Context) juggle.IModel {
	//return &models.UserModel{
	//	UserId:   101,
	//	UserName: "zhangsan",
	//}

	// 从数据库中查找数据(使用Gorm)
	user := models.NewUserModel()
	err := ctx.ShouldBindUri(user)
	//log.Println(user.UserId)
	if err != nil {
		juggle.Error(err)
	}
	this.DB.Debug().Table("user").Where("id=?",user.UserId).Find(user)
	//log.Println(user)
	return user
}

func (this *UsersController) GetUserList(ctx *gin.Context) juggle.Models {
	juggle.Error(fmt.Errorf("abc"),"err test")
	users := []*models.UserModel{
		&models.UserModel{
			UserId:   101,
			UserName: "zhangsan",
		},
		&models.UserModel{
			UserId:   101,
			UserName: "zhangsan",
		},
	}
	return juggle.MakeModels(users)
}

func (this *UsersController) Build(juggle *juggle.Juggle)  {
	juggle.Handle("GET","/users/:id",this.GetUserDetail)
	juggle.Handle("GET","/users",this.GetUserList)
}
