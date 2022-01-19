package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Module struct {
	ID   int
	Name string
}

func getModule(db *sql.DB, ID int) Module {
	modulequery := fmt.Sprintf("SELECT ModuleID, Name FROM Module WHERE ModuleID = '%d'", ID)

	moduleResults, err := db.Query(modulequery)
	if err != nil {
		panic(err.Error())
	}
	var module Module
	for moduleResults.Next() {
		moduleResults.Scan(&module.ID, &module.Name)
	}

	return module

}

func getModules(db *sql.DB) []Module {
	modulequery := "SELECT ModuleID, Name FROM Module"

	moduleResults, err := db.Query(modulequery)
	if err != nil {
		panic(err.Error())
	}

	var moduleList []Module
	for moduleResults.Next() {
		var module Module
		moduleResults.Scan(&module.ID, &module.Name)
		moduleList = append(moduleList, module)
	}
	fmt.Println(moduleList)
	return moduleList
}

func modules(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ETIAssignment2TestDB")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			moduleList := getModules(db)
			if len(moduleList) > 0 {
				fmt.Println(moduleList)
				json.NewEncoder(w).Encode(moduleList)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
	}
}

func module(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ETIAssignment2TestDB")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	moduleID := params["moduleid"]
	moduleIDint, err := strconv.Atoi(moduleID)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			moduleObject := getModule(db, moduleIDint)
			fmt.Println(moduleObject)
			json.NewEncoder(w).Encode(moduleObject)
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
	// This is t allow the headers, origins and methods all to access CORS resourcesaring
	headers := handlers.AllowedHeaders([]string{"X-Reqested-With", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	router := mux.NewRouter()

	router.HandleFunc("/api/test", testcode).Methods("GET")

	router.HandleFunc("/api/module", modules).Methods("GET")

	router.HandleFunc("/api/getmodule/{moduleid}", module).Methods("GET")

	// router.HandleFunc("/api/Rating/student/sent/{CreatorID}", postedRatings).Methods("GET")

	// router.HandleFunc("/api/Rating/class/sent/{CreatorID}", postedRatings).Methods("GET")

	// router.HandleFunc("/api/Rating/module/sent/{CreatorID}", postedRatings).Methods("GET")

	// router.HandleFunc("/api/Rating/tutor/sent/{CreatorID}", postedRatings).Methods("GET")

	// router.HandleFunc("/api/Rating/received/{CreatorID}", receivedRatings).Methods("GET")

	fmt.Println("Listening at port 5005")
	log.Fatal(http.ListenAndServe(":5005", handlers.CORS(headers, origins, methods)(router)))
}
