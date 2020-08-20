package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
)

type UserMiddleware struct {

}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (this *UserMiddleware) Onquest(ctx *gin.Context) error {
	log.Println("用户中间件")
	log.Println(ctx.Query("name")) // XXX?name=XXX
	return nil
}
