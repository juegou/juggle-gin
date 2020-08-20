package controllers

import (
	"github.com/gin-gonic/gin"
	"juggle-gin/src/juggle"
)

type IndexController struct {

}

func NewIndexController() *IndexController {
	return &IndexController{}
}

// 业务
func (this *IndexController) Index(ctx *gin.Context) string {
	return "index"
}

// 路由构建
func (this *IndexController) Build(juggle *juggle.Juggle)  {
	juggle.Handle("GET","/",this.Index)
}
