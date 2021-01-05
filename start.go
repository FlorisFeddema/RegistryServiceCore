package main

import (
	"CoreService/src/server"
	"CoreService/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.New()
	server.InitMiddleware(r)

	s := &http.Server{
		Addr: ":" + util.GetConfig().Port,
		Handler: r,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	_ = s.ListenAndServe()
}
