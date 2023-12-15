package main

import (
	"fmt"
	"time"

	log "github.com/eric-tech01/simple-log"
	"github.com/eric-tech01/taurus"
	"github.com/eric-tech01/taurus/server"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	taurus.Application
}

func main() {
	eng := &Engine{}
	if err := eng.Startup(
		eng.printLog,
		eng.serverHttp,
	); err != nil {
		panic(err)
	}
	eng.Run()
}
func (eng *Engine) serverHttp() error {
	fmt.Println("start server http")

	s := server.New()
	diag := s.Group("/taurus/diagnoise")
	diag.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, "normal")
	})
	diag.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(200, "normal")
	})

	return eng.Serve(*s)
}

func (eng *Engine) printLog() error {
	go func() {
		for {

			log.Infof("info ....")
			log.Debugf("debuf ....")
			log.Warnf("warn ....")
			log.Errorf("error ....")
			time.Sleep(3 * time.Second)
		}
	}()
	return nil
}
