package main

import (
	"fmt"
	"log"
	"net/http"
	"Lauzhack2019/backend/search"
)

func main() {
	serverAddr := "localhost:8080"
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/search", searchHandler)

	for {
		if err := http.ListenAndServe(serverAddr, nil); err != nil{
			log.Println(err)
		}

	}
}

func searchHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method{
	case "GET":
		search.Query{}
		if firstName, ok := request.URL.Query()["firstName"]; ok{
			fmt.Println(firstName)
		}


	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}