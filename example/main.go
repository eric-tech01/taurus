package main

import (
	"fmt"
	"time"

	sjson "github.com/eric-tech01/simple-json"
	log "github.com/eric-tech01/simple-log"
	version "github.com/eric-tech01/simple-version"
	"github.com/eric-tech01/taurus"
	conf "github.com/eric-tech01/taurus/pkg/conf"
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

type PostStruct struct {
	Level int `json:"level"`
}

func (eng *Engine) serverHttp() error {
	fmt.Println("start server http")
	s := server.New()

	diag := s.Group("/taurus/diagnoise")

	diag.POST("/setLevel", func(ctx *gin.Context) {
		l := &PostStruct{}
		if err := ctx.Bind(l); err != nil {
			log.Error(err)
			return
		}
		conf.Set("taurus_log_default.Level", "info")
		conf.Flush()
		rsp := sjson.New()
		rsp.Set("level", l.Level)
		log.Infof("set level %d", l.Level)
		ctx.JSON(200, rsp)
	})

	diag.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, "normal")
	})
	diag.GET("/version", func(ctx *gin.Context) {
		v := version.GetVersion()
		ctx.JSON(200, v)
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
