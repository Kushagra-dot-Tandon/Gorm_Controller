package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func retrive_id() {
	resp, err := http.Get("http://localhost:8080/1")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	fmt.Println(sb)
}

func read_all_database() {
	resp, err := http.Get("http://localhost:8080/read_all_database")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	fmt.Println(sb)

}

func main() {
	retrive_id()
	read_all_database()
}
