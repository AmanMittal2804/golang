package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //alias
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {
	encodeJson()
}

func encodeJson() {
	courses := []course{
		{"ReactJs Bootcamp", 299, "cineplix.in", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 299, "cineplix.in", "ang123", []string{"full stack", "js"}},
		{"Angular Bootcamp", 299, "cineplix.in", "aye123", []string{"web-dev", "js"}},
	}

	// finalJson, _ := json.Marshal(courses)
	finalJson, _ := json.MarshalIndent(courses, "", "\t")
	//implementation of json is Marshal

	fmt.Printf("%s\n", finalJson)
}
