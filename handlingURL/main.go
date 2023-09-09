package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ijjiu4j4riu4jiurjuj"

func main() {
	//parsing the url

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query() //key va;ue pair

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsofUrl := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
	}
	anotherURL := partsofUrl.String()
	fmt.Println(anotherURL)
}
