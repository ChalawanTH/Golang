package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Course struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Instructor string  `json:"instructor"`
}

var CourseList []Course

func init() {
	CourseJSON := `[
	{
		"id":1,
		"name":"Python",
		"price":2590,
		"instructor":"Dev-Pasit"
	},
		{
		"id":2,
		"name":"JavaScript",
		"price":0,
		"instructor":"Dev-Kiattisak"
	},
		{
		"id":3,
		"name":"SQL",
		"price":0,
		"instructor":"Dev-Pong"
	}
	]`

	err := json.Unmarshal([]byte(CourseJSON), &CourseList)
	if err != nil {
		log.Fatal(err)
	}
}

func findID(ID int) (*Course, int) {
	for i, course := range CourseList {
		if course.ID == ID {
			return &course, i
		}
	}
	return nil, 0
}
func coursesHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(CourseList)

	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPost:
		var newCourse Course
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newCourse) // Pass the address of newCourse
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		// ... (Add validation for newCourse here) ...
		if newCourse.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newCourse.ID = getNextID()
		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}
func courseHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "course/") // Correct the typo

	if len(urlPathSegment) < 2 { // Check if an ID is provided
		http.Error(w, "Missing course ID", http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	course, linstItemIndex := findID(ID)
	if course == nil {
		http.Error(w, fmt.Sprintf("no course with id %d", ID), http.StatusNotFound)
	}
	switch r.Method {
	case http.MethodGet:
		courseJSON, err := json.Marshal(course)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPut:
		var updatedCourse Course
		byteBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(byteBody, &updatedCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedCourse.ID != ID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		course = &updatedCourse
		CourseList[linstItemIndex] = *course
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getNextID() int {
	highestID := -1
	for _, course := range CourseList { // Use CourseList, not courseList
		if highestID < course.ID {
			highestID = course.ID
		}
	}
	return highestID + 1 // Return the next available ID
}
func main() {
	http.HandleFunc("/course/", courseHandler)
	http.HandleFunc("/course", coursesHandler)
	http.ListenAndServe(":5000", nil)
}
