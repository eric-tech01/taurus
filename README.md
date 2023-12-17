

## Introduction
Simplify and Accelerate Your Development Process.
Taurus is designed to provide developers with a simplified and efficient solution for building HTTP services and leveraging various utility libraries. With our framework, you can embark on rapid development journeys, ensuring both simplicity and efficiency throughout the entire process.

     _(_____)_
    |__.___.__|
        | |    
        | | 
        |_|       
  
## Documentation
To be updated...

## Requirements
- Go version >= 1.18

## Quick Start
 1.  Config file:
```
[taurus.server.http]
    Host = "0.0.0.0"
    Port = 8090
[taurus.log.default]
    Level = "debug"
    FileName = "./taurus.log"
    MaxBackups = 1
    MaxSizeInMB = 10 #dd
    Compress = true
    LocalTime = true
```

2. Example code:
```
package main

import (
	"fmt"
	"time"
	log "github.com/eric-tech01/simple-log"
	sjson "github.com/eric-tech01/simple-json"
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
		rsp := sjson.New()
		rsp.Set("level", l.Level)
		log.Infof("set level %d", l.Level)
		ctx.JSON(200, rsp)
	})

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
```

3. Run code
```
 go run ./main.go --config=config.toml
```

## Bugs and Feedback
