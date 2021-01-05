package server

import (
	"CoreService/src/util"
	"fmt"
	"github.com/getsentry/sentry-go"
	sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(router *gin.Engine)  {
	setRecovery(router)
	setLogger(router)
	setCors(router)
	setSentry(router)
	setRoutes(router)
}

func setRecovery(router *gin.Engine)  {
	router.Use(gin.Recovery())
}

func setLogger(router *gin.Engine)  {
	router.Use(gin.Logger())
}

func setCors(router *gin.Engine)  {
//	TODO add more options
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"PUT", "PATCH", "POST", "GET", "DELETE"}
	corsConfig.AllowAllOrigins = true

	router.Use(cors.New(corsConfig))
}

func setSentry(router *gin.Engine) {
	dsn := util.GetConfig().Dsn
	if len(dsn) == 0 {
		return
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "",
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	router.Use(sentryGin.New(sentryGin.Options{
		Repanic: true,
	}))
}

func setRoutes(router *gin.Engine) {

}