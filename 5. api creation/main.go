package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	// mux is a router. It is used to match the incoming request URL with a list of registered routes and calls the associated handler for that route. It is used to build REST API services.
	// goriila/mux is a powerful URL router and dispatcher for golang. example of router: http://localhost:4000/course/1
	"github.com/gorilla/mux" // go get -u github.com/gorilla/mux
)

type Course struct {
	CourseId string `json:"courseid"` // json:"courseid" is called tag
	CourseName string `json:"coursename"`
	Price int `json:"price"`
	Author *Author `json:"author"` // Author is a struct
}

type Author struct {
	FullName string `json:"fullname"`
	Website string `json:"website"`
}

// fake db
var courses []Course

// middleware/ helper function
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("learning json")

	r := mux.NewRouter() // create a new router. it is from gorilla/mux. other mux can be used. but gorilla/mux is the most popular one.

	// seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "Go", Price: 100, Author: &Author{FullName: "John Doe", Website: "johndoe.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Python", Price: 200, Author: &Author{FullName: "heyly", Website: "heyly.com"}})

	// routes. these are endpoints. example of endpoint: http://localhost:4000/course/1. by this url, one get one course. endpoint is the url that is used to access the api.
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}


// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to Api</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// convert courses to json. this way all courses will be returned.
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r) // mux.Vars(r) returns a map of route variables that are set for the current request. example: map[id:1]

	// loop through courses
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course) // return course in json format
			return
		}
	}

	// if no course found
	json.NewEncoder(w).Encode("no course found")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // decode json to course struct

	// check if course is empty
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("course is empty")
		return
	}

	// generate unique id
	//rand.Seed(time.Now().UnixNano()) // generate random number. time.Now().UnixNano() is used to generate random number. it is used to generate unique id.it is unique because it is based on time. but it is not good to use it as unique id. because it is not guaranteed to be unique. so, it is better to use uuid. uuid is a package that is used to generate unique id. it is guaranteed to be unique.

	course.CourseId = strconv.Itoa(rand.Intn(100)) // convert int to string. rand.Intn(100) is used to generate random number between 0 to 100.

	// append course to courses
	courses = append(courses, course)

	// return course
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // decode json to course struct

	// check if course is empty
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("course is empty")
		return
	}

	// grab id from request
	params := mux.Vars(r) // mux.Vars(r) returns a map of route variables that are set for the current request. example: map[id:1]

	// loop through courses
	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // delete course. append(courses[:index], courses[index+1:]...) is used to delete course. it is called slice trick. it works like this: courses[:index] is used to get all items before index. courses[index+1:] is used to get all items after index. so, courses[:index] + courses[index+1:] is used to get all items except index. so, it is used to delete item from slice.

			// update course
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// if no course found
	json.NewEncoder(w).Encode("no course found")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses
	for index, item := range courses {
		if item.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

	// if no course found
	json.NewEncoder(w).Encode("no course found")
	return
}