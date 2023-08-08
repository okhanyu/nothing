package main

import (
	"flag"
	blogconfig "nothing/config/blog"
	"nothing/internal/app/blog"
)

func main() {
	configFile := flag.String("conf", "config/blog/config.yml", "config file")
	flag.Parse()
	conf := blogconfig.ParseConfig(*configFile)
	server, err := blog.CreateBlogServer("blog", conf)
	if err != nil {
		return
	}
	server.Start(conf.System.Port)
}
