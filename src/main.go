package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"newcomic.info/api"
	"newcomic.info/db"
	"newcomic.info/log"
)

func main() {
	e := db.Init()
	if e != nil {
		panic(e)
	}
	defer db.Close()

	dataPath := os.Args[1]
	imagePath := dataPath + "/image"

	ec := echo.New()
	g := ec.Group("/api")

	ec.Static("/", "./static/")
	ec.Static("/image", imagePath)

	g.GET("/comics/:page", api.GetComicInfos)
	g.GET("/comic/:id", api.GetComicDetail)
	g.POST("/comic/favorite/:id", api.AddFavorite)
	g.DELETE("/comic/favorite/:id", api.DeleteFavorite)
	g.POST("/comic/download/:id", api.AddDownload)
	g.DELETE("/comic/download/:id", api.DeleteDownload)

	startServer(ec, ":8080")
}

func startServer(ec *echo.Echo, addr string) {
	go func() {
		log.I("start server")
		if e := ec.Start(addr); e != nil {
			if e != http.ErrServerClosed {
				log.F("start server failed, error:", e)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown with 10 seconds timeout
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.I("shutdown server ...")
	if e := ec.Shutdown(ctx); e != nil {
		log.F("shutdown server error:", e)
	}
	log.I("server closed")
}
