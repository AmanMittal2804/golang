package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	CourseId    string
	CourseName  string
	CoursePrice int
	Author      *Author
}
type Author struct {
	Fullname string
	Website  string
}

// fakedb
var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

// r is from where we are getting the response
// w is basically sending the response
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> its me here aman mittal</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-type", "apllication/json")
	json.NewEncoder(w).Encode(courses)
}

//SEEDING .!!
