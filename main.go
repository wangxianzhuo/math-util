package main

import (
	"github.com/wangxianzhuo/math-util/config"
	"github.com/wangxianzhuo/math-util/server"
)

func main() {
	conf := server.Conf{}
	config.Load(&conf, "./config.json")
	server.StartServer(conf)
}
