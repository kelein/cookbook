package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"

	"github.com/kelein/cookbook/gokite/kitgen/service/echoservice"
)

func main() {
	serv := echoservice.NewServer(
		new(EchoServiceFacade),
		server.WithServiceAddr(&net.TCPAddr{
			IP:   net.ParseIP("0.0.0.0"),
			Port: 9000,
		}),
	)

	if err := serv.Run(); err != nil {
		log.Println(err.Error())
	}
}
