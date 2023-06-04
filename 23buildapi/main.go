package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CouseId     string  `json:"courseid"`
	CourseName  string  `json:"cousename"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// Middelware/Helper
func (c *Course) IsEmpty() bool {
	// return c.CouseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

}

// controllers -files

// home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hey setting up backend API's</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Course By id")
	w.Header().Set("Content-Type", "application/json")

	// grab id form param
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CouseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id.")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	// What if request body is empty ?

	if r.Body == nil {
		json.NewEncoder(w).Encode("Request body cannot be empty.")
		return
	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Course name is required.")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CouseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Request body cannot be empty.")
		return
	}
	params := mux.Vars(r)
	// Decode JSON

	var newCourse Course
	json.NewDecoder(r.Body).Decode(&newCourse)

	if newCourse.IsEmpty() {
		json.NewEncoder(w).Encode("Course name is required.")
		return
	}

	// Find course by id and update it
	for index, course := range courses {
		if course.CouseId == params["id"] {
			courses = append(courses, newCourse)
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(newCourse)
			return
		}
	}

	json.NewEncoder(w).Encode("Course not found.")

}

func removeOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for idx, course := range courses {
		if course.CouseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			break
		}
	}
}
