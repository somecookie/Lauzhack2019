package main

import (
	"github.com/somecookie/Peerster/helper"
	"net/http"
)

func main() {
	serverAddr := "localhost:8080"
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	//http.HandleFunc("/id", IDHandler)

	for {
		err := http.ListenAndServe(serverAddr, nil)
		helper.LogError(err)
	}
}