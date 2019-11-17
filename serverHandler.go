package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/somecookie/Lauzhack2019/backend/blockchain"
	"github.com/somecookie/Lauzhack2019/backend/database"
	"github.com/somecookie/Lauzhack2019/backend/search"
	"log"
	"net/http"
	"strconv"
)

var hasher = sha256.New()
var transactions = make([]*types.Transaction,0)

func launchServer() {
	database.PopulateUsers()
	http.Handle("/", http.FileServer(http.Dir("frontend")))
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/validate", validateHandler)

	for {
		if err := http.ListenAndServe("localhost:8080", nil); err != nil {
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
		err := request.ParseForm()
		if err == nil {
			nameDoctor := request.Form.Get("nameDoctor")
			//namePatient := request.Form.Get("namePatient")
			success := request.Form.Get("success")
			successBool := success == "true"

			reportToHash := request.Form.Get("toHash")

			hasher.Reset()
			_,err := hasher.Write([]byte(reportToHash))

			if err != nil{
				log.Println(err)
				return
			}

			hash :=hasher.Sum(nil)

			tx, err := Session.AppendValidated(nameDoctor,stringToKeccak256(hex.EncodeToString(hash)),successBool)

			if err != nil{
				log.Println(err)
				return
			}

			transactions = append(transactions, tx)

	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}

func validateHandler(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	switch request.Method{
	case "POST":
		err := request.ParseForm()
		if err == nil {
			hash := request.Form.Get("hash")

			fmt.Println(hash)
		}

	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}


func loginHandler(writer http.ResponseWriter, request *http.Request) {
	enableCors(&writer)
	switch request.Method{
	case "GET":
		err := request.ParseForm()
		if err == nil {
			username := request.Form.Get("username")
			password := request.Form.Get("password")

			accepted := database.CheckUsers(username, password)

			if blockListJson, err := json.Marshal(accepted); err == nil {
				writer.Header().Set("Content-Type", "application/json")

				writer.WriteHeader(http.StatusOK)
				writer.Write(blockListJson)
			} else {
				writer.WriteHeader(http.StatusInternalServerError)
			}
		}

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

			writer.WriteHeader(http.StatusOK)
			writer.Write(blockListJson)
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
