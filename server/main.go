package server

import (
	"fmt"
	"net"
	"time"

	helper "chat-app-grpc/helper"

	homefeed "chat-app-grpc/app/homefeed"
	config "chat-app-grpc/config"
	middleware "chat-app-grpc/middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// StartServer Main rpc server Initialize
func StartServer() {
	port := config.Cfg.Ports.RpcPort
	ip := config.Cfg.Ports.RpcIp
	// ip := "192.168.159.1"
	tcp := ip + ":" + port
	lis, err := net.Listen("tcp", tcp)
	fmt.Println("Listening on", port)
	helper.SugarObj.Info("Listening on", port)

	if err != nil {
		helper.SugarObj.Error(err)
	}
	// s := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryInterceptor))
	s := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryInterceptor), grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute, // <--- This fixes it!
	}))

	//helloworld.RegisterYourServiceServer(s, &helloworld.HelloWorldStruct{})

	homefeed.RegisterHomeFeedServiceServer(s, &homefeed.HomePageStruct{})

	if err := s.Serve(lis); err != nil {
		helper.SugarObj.Error(err)

	}
}
