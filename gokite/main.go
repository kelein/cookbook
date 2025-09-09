package main

import (
	"log"

	"github.com/kelein/cookbook/gokite/kitgen/service/echoservice"
)

func main() {
	svr := echoservice.NewServer(new(EchoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
