package main

import (
	"flag"
)

func init() {
	flag.StringVar(&API_SERVERPORT, "server", "127.0.0.1:10777", "api server port")
	flag.StringVar(&CONFIG_FILE, "conf", "./service.conf", "config file")
	flag.StringVar(&DATASTORE_DIR, "datastore", "./datastore/", "datastore dir")
	flag.Parse()
}
