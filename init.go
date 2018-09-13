package main

import (
	"flag"
)

func init() {
	flag.StringVar(&API_PROXYPORT, "port", "10777", "api server port")
	flag.StringVar(&API_PROXYBIND, "bind", "127.0.0.1", "api server listen")
	flag.StringVar(&API_URL, "url", "http://test/api/v1/", "the api target")
	flag.StringVar(&API_TOKEN, "token", "token", "the api token to authenticate")
	flag.Parse()
}
