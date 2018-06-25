package main

import (
	"fmt"
	"log"
	"time"
	"server_event/serve"
	"net/http"
	"server_event/config"
)



func main() {

	broker := serve.NewServer()

	go func() {
		for {
			time.Sleep(time.Second * 5)
			eventString := fmt.Sprintf("the time is %v", time.Now())
			log.Println("Receiving event")
			broker.Notifier <- []byte(eventString)
		}
	}()
	conf := config.GetConfig()
	if err := http.ListenAndServe(conf.Bind, broker); err != nil {
		log.Fatalf("HTTP server error: %v\n", err)
	}
}
