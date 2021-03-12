package main

import (
	"CoreService/src/registry"
	"CoreService/src/repository"
	"CoreService/src/server"
	"CoreService/src/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strconv"
	"time"
)

var g errgroup.Group

func main() {
	util.SetupLogger()
	setMode()
	g.Go(func() error {
		r := gin.New()
		server.InitMiddleware(r)

		s := &http.Server{
			Addr: ":" + strconv.Itoa(util.GetConfig().Server.Http.Port),
			Handler: r,
			ReadTimeout: 10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		return s.ListenAndServe()
	})

	g.Go(func() error {
		repository.SetupConnection()
		repository.CreateDatabases()
		repository.CreateRegistry()
		registry.Test()
		return nil
	})

	if err := g.Wait(); err != nil {
		util.Logger().Fatal(err.Error())
	}
}

func setMode() {
	mode := util.GetConfig().Server.Http.Mode
	if mode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else if mode == "test" {
		gin.SetMode(gin.TestMode)
 	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	util.Logger().Info("Http server is running in " + gin.Mode() + " mode")
}
