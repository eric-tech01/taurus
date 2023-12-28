package taurus

import (
	"log"

	logger "github.com/eric-tech01/simple-log"
	"github.com/eric-tech01/taurus/pkg/flag"
	"github.com/eric-tech01/taurus/utils"
	"github.com/sirupsen/logrus"

	conf "github.com/eric-tech01/taurus/pkg/conf"
)

func init() {

	flag.Register(&flag.StringFlag{Name: "config", Usage: "--config=config.toml", Default: "config.toml", Action: func(key string, fs *flag.FlagSet) {
		configAddr := fs.String(key)
		log.Printf("read config: %s", configAddr)
		err := conf.Load(configAddr)
		if err != nil {
			log.Fatalf("build datasource[%s] failed: %v", configAddr, err)
		}
	}})
	flag.Parse()

	//register log hook
	initLogger()
}

func initLogger() {
	conf.OnLoaded(func(c *conf.Configuration) {
		log.Println("load logger start")
		defer log.Println("load logger finish")
		m := conf.GetStringMap("taurus.log.default")
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
		op := &logger.Option{}
		utils.MapToStruct(m, op)
		logger.SetOptions(fileName, op)
		logger.SetLevel(level)
	})
}
