package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "67533", CourseName: "Golang", CoursePrice: 998, Author: &Author{Fullname: "amanMittal", Website: "mittalaman284@gmail.com"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
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

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one courses")
	w.Header().Set("Content-type", "apllication/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No couse found")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one courses")
	w.Header().Set("Content-type", "apllication/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send dome data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//GENERATE UNIQUE ID
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	return
}
