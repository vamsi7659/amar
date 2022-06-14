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
	name     string  `json:"name"`
	CourseId string  `json:"CourseId"`
	price    int     `json:price"`
	writer   *writer `json:"writer"`
}
type writer struct {
	fullname string `json:"author"`
	website  string `json:"age"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	//return c.name =="" && c.course =
	return c.name == ""
}

func main() {
	fmt.Println(" kjhfr general api project")
	r := mux.NewRouter()
	courses = append(courses, Course{name: "golang", CourseId: "1",
		price: 20, writer: &writer{fullname: "vamsi", website: "vamsi.com"}})
	courses = append(courses, Course{name: "devops", CourseId: "2",
		price: 20, writer: &writer{fullname: "varun", website: "varun.com"}})
	r.HandleFunc("/", servehome).Methods("GET")
	r.HandleFunc("/courses", getallcourses).Methods("GET")
	r.HandleFunc("/courses/{Id}", getonecourse).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", r))

}

func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to api<"))
}
func getallcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("contenttype", "appliction/json")
	json.NewEncoder(w).Encode(courses)
}
func getonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("contenttype", "appliction/json")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["Id"] {
			json.NewEncoder(w).Encode(courses)
			return
		}
		json.NewEncoder(w).Encode("no course found")
		return
	}

}
func addingcourse(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		json.NewEncoder(w).Encode("no course found")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("send data")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
	return

}
