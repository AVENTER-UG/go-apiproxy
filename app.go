package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	_ "net/http/pprof"
	"net/url"
)

var API_PROXYPORT string
var API_PROXYBIND string
var API_URL string
var API_TOKEN string

var MinVersion string

var srv http.Server

type handle struct {
	reverseProxy string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(this.reverseProxy + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
	remote, err := url.Parse(this.reverseProxy)
	if err != nil {
		log.Fatalln(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	r.Host = remote.Host
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %v", API_TOKEN))
	proxy.ServeHTTP(w, r)
}

func main() {
	log.Println("GO-APIPROXY build"+MinVersion, API_PROXYBIND, API_PROXYPORT, API_URL)
	srv.Handler = &handle{reverseProxy: API_URL}
	srv.Addr = API_PROXYBIND + ":" + API_PROXYPORT
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
