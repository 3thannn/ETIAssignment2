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

type Tutor struct {
	ID   string
	Name string
}

func getTutors(db *sql.DB) []Tutor {
	tutorQuery := "SELECT TutorID, Name FROM Tutor"

	tutorResults, err := db.Query(tutorQuery)
	if err != nil {
		panic(err.Error())
	}

	var tutorList []Tutor
	for tutorResults.Next() {
		var tutor Tutor
		tutorResults.Scan(&tutor.ID, &tutor.Name)
		tutorList = append(tutorList, tutor)
	}

	return tutorList
}

func getTutor(db *sql.DB, ID string) Tutor {
	tutorQuery := fmt.Sprintf("SELECT TutorID, Name FROM Tutor WHERE TutorID = '%s'", ID)

	tutorResults, err := db.Query(tutorQuery)
	if err != nil {
		panic(err.Error())
	}

	var tutor Tutor
	for tutorResults.Next() {
		tutorResults.Scan(&tutor.ID, &tutor.Name)
	}

	return tutor
}

func tutors(w http.ResponseWriter, r *http.Request) {
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
			tutorList := getTutors(db)
			if len(tutorList) > 0 {
				fmt.Println(tutorList)
				json.NewEncoder(w).Encode(tutorList)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
	}
}

func tutor(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(db3:9049)/ETIAssignment2TestDB")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	tutorID := params["tutorid"]
	if r.Method == "GET" {
		if err == nil {
			tutor := getTutor(db, tutorID)
			fmt.Println(tutor)
			json.NewEncoder(w).Encode(tutor)
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

	router.HandleFunc("/api/tutor", tutors).Methods("GET")

	router.HandleFunc("/api/gettutor/{tutorid}", tutor).Methods("GET")

	fmt.Println("Listening at port 9044")
	log.Fatal(http.ListenAndServe(":9044", handlers.CORS(headers, origins, methods)(router)))
}
