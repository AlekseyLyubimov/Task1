package main

import (
	"log"
	"net/http"
	"os"
)

func handleRequests() {
	http.HandleFunc("/api/v1/xml", returnXml)
	http.HandleFunc("/api/v1/json", returnJson)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	log.Println("Application start")
	handleRequests()
}

func returnXml(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnXml")
	w.Write(readFromFile("return.xml"))
}

func returnJson(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnJson")
	w.Write(readFromFile("return.json"))
}

func readFromFile(fileName string) []byte {
	b, err := os.ReadFile(fileName) // just pass the file name
	if err != nil {
		log.Println(err)
	}
	return b
}
