package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"

	"hackerthon2019/redis-ha/redis-ha-demo/app"
)

const (
	port = "8888"
)

func run(gs *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: gs,
	}

	go func(srv *http.Server) {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln("listen and serve:", err.Error())
		}
	}(srv)

	return srv
}

func stop(srv *http.Server) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	srv.Shutdown(ctx)
}

func main() {
	gin.SetMode(gin.DebugMode)

	gs := gin.Default()
	gs.LoadHTMLGlob("static/*")
	gs.Static("/static", "static")

	app.Init(gs)

	srv := run(gs)

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	signal.Notify(interrupt, os.Kill)
	<-interrupt

	stop(srv)
}
