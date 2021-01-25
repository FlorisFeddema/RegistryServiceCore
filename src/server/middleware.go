package server

import (
	"CoreService/src/util"
	"github.com/getsentry/sentry-go"
	sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(router *gin.Engine)  {
	setLogger(router)
	setCors(router)
	setSentry(router)
	setRoutes(router)
}


func setLogger(router *gin.Engine)  {
	router.Use(ginzap.Ginzap(util.Logger(), time.RFC3339, true))
	router.Use(gin.Recovery())
}

func setCors(router *gin.Engine)  {
//	TODO add more options
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"PUT", "PATCH", "POST", "GET", "DELETE"}
	corsConfig.AllowAllOrigins = true

	router.Use(cors.New(corsConfig))
}

func setSentry(router *gin.Engine) {
	dsn := util.GetConfig().Sentry.Dsn
	if len(dsn) == 0 {
		return
	}

	util.Logger().Info("Setting up connection with Sentry")

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "",
	}); err != nil {
		util.Logger().Fatal(err.Error())
	}

	router.Use(sentryGin.New(sentryGin.Options{
		Repanic: true,
	}))
}

func setRoutes(router *gin.Engine) {

}