package main

import "os"

func init() {
	API_TOKEN = os.Getenv("API_TOKEN")
	API_PROXYBIND = os.Getenv("API_PROXYBIND")
	API_PROXYPORT = os.Getenv("API_PROXYPORT")
	API_URL = os.Getenv("API_URL")
}
