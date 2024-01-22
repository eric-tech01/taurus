package taurus

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	logger "github.com/eric-tech01/simple-log"
	"github.com/eric-tech01/taurus/pkg/flag"
	"github.com/eric-tech01/taurus/utils"
	"github.com/sirupsen/logrus"

	conf "github.com/eric-tech01/taurus/pkg/conf"
)

func init() {
	//register log hook
	initLogger()
	flag.Register(&flag.StringFlag{Name: "config", Usage: "--config=config.toml", Default: "config.toml", Action: func(key string, fs *flag.FlagSet) {
		configAddr := fs.String(key)
		log.Printf("read config: %s", configAddr)
		err := conf.Load(configAddr)
		if err != nil {
			log.Fatalf("build datasource[%s] failed: %v", configAddr, err)
		}
	}})
	flag.Parse()
}

func initLogger() {
	conf.OnLoaded(func(c *conf.Configuration) {
		log.Println("load logger start")
		defer log.Println("load logger finish")
		m := conf.GetStringMap("taurus_log_default")
		if len(m) == 0 {
			log.Printf("has no log config, ignore load")
			return
		}
		var fileName string
		if f, ok := m["FileName"]; ok {
			fileName = f.(string)
		}
		var level logrus.Level = logrus.DebugLevel
		if l, ok := m["Level"]; ok {
			level, _ = logrus.ParseLevel(l.(string))
		}
		op := &logger.Option{Formatter: &myFormatter{}}
		utils.MapToStruct(m, op)
		logger.SetOptions(fileName, op)
		logger.SetLevel(level)
	})
}

// myFormatter 自定义日志格式
type myFormatter struct {
}

// Format 格式化日志
func (f *myFormatter) Format(e *logrus.Entry) ([]byte, error) {
	fileName, line := getFileLine(9)
	return []byte(fmt.Sprintf("%s %5.5s [%s:%v %s] %s\n",
		e.Time.Local().Format("2006/01/02 15:04:05.000000"),
		e.Level.String(),
		fileName,
		line, splitAndGetLast(e.Caller.Function, "."),
		e.Message)), nil
}

// splitAndGetLast 分割字符串并返回最后一个元素
func splitAndGetLast(str string, sep string) string {
	slice := strings.Split(str, sep)
	return slice[len(slice)-1]
}
func getFileLine(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "???"
		line = 0
	} else {
		_, file = filepath.Split(file)
	}
	return file, line
}
