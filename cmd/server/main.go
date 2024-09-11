package main

import (
	"log"

	"github.com/foreverLoveWisdom/go_distributed_proglog.git/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
