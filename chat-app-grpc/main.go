package main

import (
	config "chat-app-grpc/config"
	server "chat-app-grpc/server"
	"fmt"
)

func main() {

	fmt.Printf("%+v", config.Cfg)
	config.Cfg.DbConnect()
	defer config.CloseDb()
	server.StartServer()
}
