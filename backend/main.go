package main

import (
	"Lauzhack2019/backend/search"
	"encoding/json"
	"fmt"
	"github.com/AntonRagot/Peerster/utils"
	"log"
	"net/http"
)

func main() {
	serverAddr := "localhost:8080"
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/get", getHandler)

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

func writeHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method{
	case "POST":
		firstName := request.Form.Get("FirstName")
		name := request.Form.Get("Name")
		location := request.Form.Get("Location")
		success := request.Form.Get("Success")
		report := request.Form.Get("Report")

		//Create block
		//Push block to the blockchain

	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}

func getHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method{
	case "GET":
		//Make a list of all blocks

		blockList := getBlocksAsList()
		blockListJson, err := json.Marshal(blockList)
		utils.CheckError(err)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(blockListJson)

	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}