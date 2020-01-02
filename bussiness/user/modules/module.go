package modules

import(
	"fmt"
	"github.com/feekk/zddgo"
	mid "github.com/feekk/zddgo/middleware"
	"github.com/feekk/zddgo/response"
	"github.com/feekk/zddgo/utils"
	"github.com/gin-gonic/gin"

	"znda/bussiness/user/modules/controller"
)


func router(handler *gin.Engine){
	userHandler := handler.Group(fmt.Sprintf("/%s", zddgo.Conf.App.Name))
	userHandler.POST("/register", controller.Register)
}

func middleware(handler *gin.Engine){
	handler.Use(mid.Recovery(), mid.Route())
	handler.GET("/ping", func(ctx *gin.Context) {response.JSONSuccess(ctx, "ok")})
}

func Init() (handler *gin.Engine) {
	utils.StdPrint(`[Run_Http] Runing Run_Http Success.
 - Env Params: %+v
 - Http Options:%+v
`, zddgo.Conf.Env, zddgo.Conf.Http)

	gin.SetMode(zddgo.Conf.Env.GinMode)
	handler = gin.New()
	middleware(handler)
	router(handler)
	return
}