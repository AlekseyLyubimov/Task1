package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	w.Write(readFromFile("return.xml"))
}

func returnJson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnJson")
	w.Write(readFromFile("return.json"))
}

func readFromFile(fileName string) []byte {

	b, err := os.ReadFile(fileName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'

	fmt.Println(str) // print the content as a 'string'

	return b
}
