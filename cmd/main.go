package main

import (
	"flag"
	"github.com/braverior/log4go"
	"larkapp-example/conf"
	"larkapp-example/internal/event_handler"
	"larkapp-example/internal/larkapp"
	"log"
	"os"
	"path"
)

var __version__ string

func main() {
	pwd, _ := os.Getwd()
	log.Println("Running Version ", __version__)
	executeDir := flag.String("d", pwd, "execute directory")
	configFile := flag.String("c", "conf/server.yaml", "default conf/server.yaml")
	logConfigFilePath := flag.String("l", "conf/logging.xml", "logging config file, relative path")
	flag.Parse()
	log4go.LoadConfiguration(path.Join(*executeDir, *logConfigFilePath))
	err := conf.ConfigInit(*configFile)
	if err != nil {
		log.Fatalf("config init failed, err %v", err)
	}

	eventHandler := event_handler.NewEventHandler()

	app := larkapp.NewLarkApp(eventHandler, conf.GlobalConfig.LarkAppConfig)
	err = app.Start()
	log.Println(err)
}
