package main

import (
	"fmt"
	_ "./api"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var API_SERVERPORT string
var CONFIG_FILE string
var DATASTORE_DIR string

func main() {
	fmt.Println("API_SERVERPORT=", API_SERVERPORT)
	fmt.Println("CONFIG_FILE=", CONFIG_FILE)
	fmt.Println("DATASTORE_DIR=", DATASTORE_DIR)

	log.Println("Listening..." + API_SERVERPORT)
	http.ListenAndServe(API_SERVERPORT, nil)
}
