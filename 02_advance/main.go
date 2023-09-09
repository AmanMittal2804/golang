package main

import "fmt"

func main() {
	fmt.Println("maps in go lang")

	lang := make(map[string]string)

	lang["JS"] = "Jvascript"
	lang["RB"] = "Ruby"

	delete(lang, "RB")

	for key, value := range lang {
		fmt.Printf("For key %v , value is %v \n", key, value)
	}

	aman := User("Hitesh", "hitesh@go.dev", true)
	fmt.Println(aman)

}

type User struct {
	Name   string
	Email  string
	Status bool
}
