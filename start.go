package main

import (
	"CoreService/src/repository"
	"CoreService/src/server"
	"CoreService/src/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var g errgroup.Group

func main() {


	g.Go(func() error {
		r := gin.New()
		server.InitMiddleware(r)

		s := &http.Server{
			Addr: ":" + util.GetConfig().Port,
			Handler: r,
			ReadTimeout: 10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		return s.ListenAndServe()
	})

	g.Go(func() error {
		repository.Connect()
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
