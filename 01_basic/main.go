package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const LoginToken string = "Vfduw321" // first letter capital means public
func main() {
	fmt.Println("Hello its me")
	var user string = "aman"
	fmt.Printf("he is %T \n", user) // type define
	var another int
	fmt.Println(another) // no garbage value

	bookingId := 300000 // can't be used outside function
	fmt.Println(bookingId)

	//taking input from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter about yourself")
	input, _ := reader.ReadString('\n') // , error syntax
	fmt.Println("Thanks for rating", input)

	//string conversion
	// numRating , err := strconv.ParseFloat(input,64)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating", numRating+1)
	}

	mynumber := 23
	var ptr = &mynumber

	fmt.Println("value of poiniter ", ptr)
	fmt.Println("value of poiniter ", *ptr)

	// pointer is done on the actual values not the copies

	var fruit [4]string

	fruit[0] = "Apple"
	fruit[1] = "tomato"
	fruit[3] = "peach"

	fmt.Println("Fruit list is :", len(fruit)) // will display 4 as it does not really calculate the value

	//slices

	var fruitList = []string{"apple", "mango"}
	fmt.Printf("the list %T\n", fruitList)

	fruitList = append(fruitList, "Peach")
	fruitList = append(fruitList[1:])
	// vale given is not included

	var num1 = []int{12, 13, 14, 15}
	sort.Ints(num1)
}
