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

type Class struct {
	ID   int
	Name string
}

func getClasses(db *sql.DB) []Class {
	classQuery := "SELECT ClassID, Name FROM Class"

	classResults, err := db.Query(classQuery)
	if err != nil {
		panic(err.Error())
	}

	var classList []Class
	for classResults.Next() {
		var class Class
		classResults.Scan(&class.ID, &class.Name)
		classList = append(classList, class)
	}

	return classList
}

func class(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ETIAssignment2TestDB")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			classList := getClasses(db)
			if len(classList) > 0 {
				fmt.Println(classList)
				json.NewEncoder(w).Encode(classList)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
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

	router.HandleFunc("/api/student", class).Methods("GET")

	// router.HandleFunc("/api/Rating/student/sent/{CreatorID}", postedRatings).Methods("GET")

	// router.HandleFunc("/api/Rating/class/sent/{CreatorID}", postedRatings).Methods("GET")

	// router.HandleFunc("/api/Rating/module/sent/{CreatorID}", postedRatings).Methods("GET")

	// router.HandleFunc("/api/Rating/tutor/sent/{CreatorID}", postedRatings).Methods("GET")

	// router.HandleFunc("/api/Rating/received/{CreatorID}", receivedRatings).Methods("GET")

	fmt.Println("Listening at port 5006")
	log.Fatal(http.ListenAndServe(":5006", handlers.CORS(headers, origins, methods)(router)))
}
