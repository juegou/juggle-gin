package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"juggle-gin/app/models"
	"juggle-gin/src/juggle"
	"juggle-gin/src/juggle/adapter"
	"log"
)

type NewsController struct {
	*adapter.XormAdapter
}

func NewNewsController() *NewsController {
	return &NewsController{}
}

func (this *NewsController) Test(ctx *gin.Context) string {
	return "test"
}

func (this *NewsController) GetNewsDetail(ctx *gin.Context) juggle.IModel {
	//return &models.News{
	//	Id:    300,
	//	NewsTitle: "title",
	//}

	// 从数据库中查找数据(使用Xorm)
	news := models.NewNewsModel()
	err := ctx.ShouldBindUri(news)
	log.Println(news.Id)
	if err != nil {
		juggle.Error(err)
	}
	//this.Table("news").Where("id=?",news.Id).Get(news)
	this.Engine.ID(news.Id).Get(news)
	return news
}

func (this *NewsController) GetNewsList(ctx *gin.Context) juggle.Models {
	juggle.Error(fmt.Errorf("abc"),"err test")
	news := []*models.News{
		&models.News{
			Id:    301,
			NewsTitle: "title1",
		},
		&models.News{
			Id:    302,
			NewsTitle: "title2",
		},
	}
	return juggle.MakeModels(news)
}


func (this *NewsController) Build(juggle *juggle.Juggle)  {
	juggle.Handle("GET","news",this.GetNewsList)
	juggle.Handle("GET","news/:id",this.GetNewsDetail)
	juggle.Handle("GET","test1",this.Test)
}


