package main

import (
	"flag"
	blogconfig "nothing/config"
	"nothing/internal/app"
)

func main() {
	configFile := flag.String("conf", "conf/config.yml", "config file")
	serverName := flag.String("name", "nothing", "server name")
	flag.Parse()
	conf := blogconfig.ParseConfig(*configFile)
	server, err := app.CreateServer(*serverName, conf)
	if err != nil {
		return
	}
	server.Start(conf.System.Port)
}
