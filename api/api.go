package api

import (
	//"encoding/json"
	"github.com/gorilla/mux"
	//"io/ioutil"
	"log"
	"net/http"
	"strings"
	"encoding/base64"
)

type userStruct struct {
	ID int
	NAME string
	PRODUCT_ID int
}

func init() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/versions", api_versions).Methods("GET")
	rtr.HandleFunc("/api", api_versions).Methods("GET")
	rtr.HandleFunc("/api/v0", api_v0_version).Methods("GET")
	rtr.HandleFunc("/api/v0/version", api_v0_version).Methods("GET")

	rtr.HandleFunc("/api/v0/authUser", api_v0_authUser).Methods("GET")

	http.Handle("/", rtr)
}

func api_versions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "-")
	w.Write([]byte("/api/v0"))
}

func api_v0_version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "v0")
	w.Write([]byte("v0.1"))
}

func api_v0_authUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		http.Error(w, "Not authorized", 401)
		return
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		http.Error(w, "Not authorized", 401)
		return
	}

	//var username = pair[0]
	//var password = pair[1]

	// TODO Check Username Password

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "v0")

	d := `{ "res":1, "failure":0 }`

	log.Println(d)

	w.Write([]byte(string(d)))
}
