package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

func courseHandler(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":5000", nil)

}
