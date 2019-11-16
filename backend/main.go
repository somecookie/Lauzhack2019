package main

import (
	"Lauzhack2019/backend/blockchain"
	"Lauzhack2019/backend/search"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	serverAddr := "localhost:8080"
	http.Handle("/", http.FileServer(http.Dir("../frontend")))
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/get", getHandler)

	for {
		if err := http.ListenAndServe(serverAddr, nil); err != nil {
			log.Println(err)
		}

	}
}

func searchHandler(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	switch request.Method {
	case "GET":
		searchQuery := parseSearchParams(request)
		blocks := searchQuery.Search()

		if res, err := json.Marshal(blocks); err == nil {
			writer.WriteHeader(http.StatusOK)
			writer.Write(res)
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}

func writeHandler(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	switch request.Method{
	case "POST":
		/*firstName := request.Form.Get("FirstName")
		name := request.Form.Get("Name")
		location := request.Form.Get("Location")
		success := request.Form.Get("Success")
		report := request.Form.Get("Report")

		fmt.Println(firstName + name + location + success + report)*/

		//Create block
		//Push block to the blockchain
	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}

func getHandler(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	switch request.Method{
	case "GET":
		//Make a list of all blocks
		//blockList := getBlocksAsList()

		blockList := blockchain.GetMockDataAsList()
		//blockList := getBlocksAsList()
		if blockListJson, err := json.Marshal(blockList); err == nil {
			writer.Header().Set("Content-Type", "application/json")

			writer.Write(blockListJson)
			writer.WriteHeader(http.StatusOK)
		} else {
			writer.WriteHeader(http.StatusInternalServerError)
		}

	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func parseSearchParams(request *http.Request) *search.Query {
	searchQuery := &search.Query{}
	if firstName, ok := request.URL.Query()["firstName"]; ok {
		searchQuery.FirstName = firstName[0]
	}
	if name, ok := request.URL.Query()["name"]; ok {
		searchQuery.Name = name[0]
	}
	if location, ok := request.URL.Query()["location"]; ok {
		searchQuery.Location = location[0]
	}
	if isSuccess, ok := request.URL.Query()["isSuccess"]; ok {
		if isSuccess[0] == "1" {
			searchQuery.IsSuccess = true
		} else if isSuccess[0] == "0" {
			searchQuery.IsSuccess = false
		}
	}
	if minSuccess, ok := request.URL.Query()["minSuccess"]; ok {
		if rate, err := strconv.ParseFloat(minSuccess[0], 64); err == nil {
			if rate >= 0 && rate <= 100 {
				searchQuery.MinSuccessRate = rate
			}
		}
	}
	if reportHash, ok := request.URL.Query()["reportHash"]; ok {
		if h, err := hex.DecodeString(reportHash[0]); err == nil && len(h) == 32 {
			searchQuery.ReportHash = reportHash[0]
		}
	}

	return searchQuery
}
