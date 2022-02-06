package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Student struct {
	ID   string
	Name string
}

func getStudents(db *sql.DB) []Student {
	studentquery := "SELECT StudentID, Name FROM Student"

	studentresults, err := db.Query(studentquery)
	if err != nil {
		panic(err.Error())
	}

	var studentList []Student
	for studentresults.Next() {
		var student Student
		studentresults.Scan(&student.ID, &student.Name)
		studentList = append(studentList, student)
	}

	return studentList
}

func getStudent(db *sql.DB, ID string) Student {
	studentquery := fmt.Sprintf("SELECT StudentID, Name FROM Student WHERE StudentID = '%s'", ID)

	studentresults, err := db.Query(studentquery)
	if err != nil {
		panic(err.Error())
	}

	var student Student
	for studentresults.Next() {
		studentresults.Scan(&student.ID, &student.Name)
	}
	return student
}

func students(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(db3:9049)/ETIAssignment2TestDB")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			studentList := getStudents(db)
			if len(studentList) > 0 {
				fmt.Println(studentList)
				json.NewEncoder(w).Encode(studentList)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
	}
}

func student(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(db3:9049)/ETIAssignment2TestDB")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	studentID := params["studentid"]
	if r.Method == "GET" {
		if err == nil {
			student := getStudent(db, studentID)
			fmt.Println(student)
			json.NewEncoder(w).Encode(student)
		} else {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
	}
}

func testcode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode("Hello this is a pass")
		w.WriteHeader(http.StatusCreated)
	}
}

func main() {
	// This is to allow the headers, origins and methods all to access CORS resource sharing
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	router := mux.NewRouter()
	router.HandleFunc("/api/test", testcode).Methods("GET")

	router.HandleFunc("/api/student", students).Methods("GET")

	router.HandleFunc("/api/getstudent/{studentid}", student).Methods("GET")

	fmt.Println("Listening at port 9043")
	log.Fatal(http.ListenAndServe(":9043", handlers.CORS(headers, origins, methods)(router)))
}
