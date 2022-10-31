package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/homepage", homePage)
	http.HandleFunc("/api/v1/xml", returnXml)
	http.HandleFunc("/api/v1/json", returnJson)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("Successfull start 1")
	handleRequests()
}

func returnXml(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnXml")
	json.NewEncoder(w).Encode("this should be XML file")
}

func returnJson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnJson")
	xml.NewEncoder(w).Encode("this should be Json file")
}
