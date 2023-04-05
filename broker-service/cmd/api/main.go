package main

import (
	"log"
	"fmt"
	"net/http"
)

const webPort = "8000"

type Config struct {

}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port: %s\n", webPort)

	//defines http server
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	//start the server
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}