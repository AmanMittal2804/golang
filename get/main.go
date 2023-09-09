package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	getRequest()
	postRequest()
}

func getRequest() {
	const url = "http://localhost:8000/get"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)

	content, _ := ioutil.ReadAll(response.Body)
	//content is inbyte format
	fmt.Println(string(content))
}

func postRequest() {
	const url = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	{
       "coursename":"Golang",
	   "name":"aman"
	}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(content)
}
