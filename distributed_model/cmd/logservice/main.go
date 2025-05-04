package main

import (
	"context"
	"fmt"

	"example.com/m/log"
	"example.com/m/service"

	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done() // cancel调用时触发

	fmt.Println("Shuting down log service.")
}

/*
curl -X POST http://localhost:4000/log -d 'hello world'
*/
