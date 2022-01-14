package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Rating struct {
	RatingID          int
	CreatorID         int
	CreatorType       string
	RatingType        string
	TargetID          int
	RatingScore       int
	Anonymous         int
	DateTimePublished string
}

type Commenter struct {
	StudentID int
	Name      string
}

func getAllStudents(db *sql.DB) []Commenter {
	url := "http://localhost:5003/api/student"
	response, err := http.Get(url)
	var StudentList []Commenter
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode == http.StatusNotFound {
			fmt.Println("409 - No Students Found!")
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			response.Body.Close()
			json.Unmarshal(data, &StudentList)
			fmt.Println("202 - Successfully received Students!")
		}
	}
	return StudentList
}

func getStudentRatings(db *sql.DB, CreatorID int) []Rating {
	studentquery := fmt.Sprintf("SELECT * FROM Rating WHERE CreatorID = '%i' AND RatingType = Student", CreatorID)

	studentresults, err := db.Query(studentquery)
	if err != nil {
		panic(err.Error())
	}

	var studentRatingList []Rating
	for studentresults.Next() {
		var Rating Rating
		Rating.RatingType = "Student"
		studentresults.Scan(&Rating.RatingID, &Rating.CreatorID, &Rating.CreatorType, &Rating.TargetID, &Rating.RatingType, &Rating.RatingScore, &Rating.Anonymous, Rating.DateTimePublished)
		studentRatingList = append(studentRatingList, Rating)
	}

	return studentRatingList
}

//STILL NEED TO DO GET STUDENT RatingS FROM TUTOR

func getClassRatings(db *sql.DB, CreatorID int) []Rating {
	classquery := fmt.Sprintf("SELECT * FROM Rating WHERE CreatorID = '%i' AND RatingType = Class", CreatorID)
	classresults, err := db.Query(classquery)
	if err != nil {
		panic(err.Error())
	}
	var classRatingList []Rating
	for classresults.Next() {
		var Rating Rating
		Rating.RatingType = "Class"
		classresults.Scan(&Rating.RatingID, &Rating.CreatorID, &Rating.CreatorType, &Rating.TargetID, &Rating.RatingType, &Rating.RatingScore, &Rating.Anonymous, Rating.DateTimePublished)
		classRatingList = append(classRatingList, Rating)
	}
	return classRatingList
}

func getModuleRatings(db *sql.DB, CreatorID int) []Rating {
	modulequery := fmt.Sprintf("SELECT * FROM Rating WHERE CreatorID = '%i' AND RatingType = Module", CreatorID)
	moduleresults, err := db.Query(modulequery)
	if err != nil {
		panic(err.Error())
	}
	var moduleRatingList []Rating
	for moduleresults.Next() {
		var Rating Rating
		Rating.RatingType = "Module"
		moduleresults.Scan(&Rating.RatingID, &Rating.CreatorID, &Rating.CreatorType, &Rating.TargetID, &Rating.RatingType, &Rating.RatingScore, &Rating.Anonymous, Rating.DateTimePublished)
		moduleRatingList = append(moduleRatingList, Rating)
	}
	return moduleRatingList
}

func getTutorRatings(db *sql.DB, CreatorID int) []Rating {
	tutorquery := fmt.Sprintf("SELECT * FROM Rating WHERE CreatorID = '%i' AND RatingType = Tutor", CreatorID)
	tutorresults, err := db.Query(tutorquery)
	if err != nil {
		panic(err.Error())
	}
	var tutorRatingList []Rating
	for tutorresults.Next() {
		var Rating Rating
		Rating.RatingType = "Tutor"
		tutorresults.Scan(&Rating.RatingID, &Rating.CreatorID, &Rating.CreatorType, &Rating.TargetID, &Rating.RatingType, &Rating.RatingScore, &Rating.Anonymous, Rating.DateTimePublished)
		tutorRatingList = append(tutorRatingList, Rating)
	}
	return tutorRatingList
}

//Get all Ratings received
func getReceivedRatings(db *sql.DB, CreatorID int) []Rating {

	query := fmt.Sprintf("SELECT * FROM Rating WHERE TargetID = '%i' AND RatingType = Student", CreatorID)

	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var RatingList []Rating
	for results.Next() {
		var Rating Rating
		results.Scan(&Rating.RatingID, &Rating.CreatorID, &Rating.CreatorType, &Rating.TargetID, &Rating.RatingType, &Rating.RatingScore, &Rating.Anonymous, Rating.DateTimePublished)
		RatingList = append(RatingList, Rating)
	}
	return RatingList
}

func postRating(db *sql.DB, Rating Rating) {
	CreatorID := Rating.CreatorID
	println(Rating.CreatorID)
	CreatorType := Rating.CreatorType
	println(Rating.CreatorType)
	TargetID := Rating.TargetID
	println(Rating.TargetID)
	RatingScore := Rating.RatingScore
	println(Rating.RatingScore)
	Anonymous := Rating.Anonymous
	println(Rating.Anonymous)
	RatingType := Rating.RatingType
	query := ""
	if Rating.RatingType == "Student" {
		query = fmt.Sprintf("INSERT INTO Rating (CreatorID, CreatorType, TargetID, RatingType, RatingScore, Anonymous) VALUES ('%d', '%s', '%d', '%s', '%d', '%b')",
			CreatorID, CreatorType, TargetID, RatingType, RatingScore, Anonymous)
	} else if Rating.RatingType == "Class" {
		query = fmt.Sprintf("INSERT INTO Rating (CreatorID, CreatorType, TargetID, RatingType, RatingScore, Anonymous) VALUES ('%d', '%s', '%d', '%s', '%d', '%b')",
			CreatorID, CreatorType, TargetID, RatingType, RatingScore, Anonymous)
	} else if Rating.RatingType == "Module" {
		query = fmt.Sprintf("INSERT INTO Rating (CreatorID, CreatorType, TargetID, RatingType, RatingScore, Anonymous) VALUES ('%d', '%s', '%d', '%s', '%d', '%b')",
			CreatorID, CreatorType, TargetID, RatingType, RatingScore, Anonymous)
	} else if Rating.RatingType == "Tutor" {
		query = fmt.Sprintf("INSERT INTO Rating (CreatorID, CreatorType, TargetID, RatingType, RatingScore, Anonymous) VALUES ('%d', '%s', '%d', '%s', '%d', '%b')",
			CreatorID, CreatorType, TargetID, RatingType, RatingScore, Anonymous)
	}
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}

func updateRecord(db *sql.DB, Rating Rating) {
	RatingID := Rating.RatingID
	RatingScore := Rating.RatingScore
	query := ""
	if Rating.RatingType == "Student" {
		query = fmt.Sprintf("UPDATE Rating SET RatingScore = '%d' WHERE RatingID = '%d' AND RatingType = Student", RatingScore, RatingID)
	} else if Rating.RatingType == "Class" {
		query = fmt.Sprintf("UPDATE  Rating SET RatingScore = '%d' WHERE RatingID = '%d' AND RatingType = Class", RatingScore, RatingID)
	} else if Rating.RatingType == "Module" {
		query = fmt.Sprintf("UPDATE Rating SET RatingScore = '%d' WHERE RatingID = '%d' AND RatingType = Module", RatingScore, RatingID)
	} else if Rating.RatingType == "Tutor" {
		query = fmt.Sprintf("UPDATE Rating SET RatingScore = '%d' WHERE RatingID = '%d' AND RatingType = Tutor", RatingScore, RatingID)
	}
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}

func rating(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ETIAssignment2Rating")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "POST" {
		var newRating Rating
		reqBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(reqBody, &newRating)
			postRating(db, newRating)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("201 - Rating Posted!"))
		}
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("422 - Please enter trip detais in JSON format!"))
	}
}

//Get all Ratings received
func receivedRatings(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ETIAssignment2Rating")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	CreatorID := params["CreatorID"]
	CreatorIDint, err := strconv.Atoi(CreatorID)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			var personalRatings []Rating = getReceivedRatings(db, CreatorIDint)
			if len(personalRatings) > 0 {
				fmt.Println(personalRatings)
				json.NewEncoder(w).Encode(personalRatings)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
	}
	if r.Method == "PUT" {
		// read the string sent to the service
		var RatingInput Rating
		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			// convert JSON to object
			json.Unmarshal(reqBody, &RatingInput)
			updateRecord(db, RatingInput)
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("202 - Your Rating has been successfully updated!"))
		}
	} else {
		w.WriteHeader(
			http.StatusUnprocessableEntity)
		w.Write([]byte("422 - Please enter account detais in JSON format!"))
	}
}

func postedStudentRatings(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ETIAssignment2Rating")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	CreatorID := params["CreatorID"]
	CreatorIDint, err := strconv.Atoi(CreatorID)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			studentRatingList := getStudentRatings(db, CreatorIDint)
			if len(studentRatingList) > 0 {
				fmt.Println(studentRatingList)
				json.NewEncoder(w).Encode(studentRatingList)
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

	router.HandleFunc("/api/rating", rating).Methods("POST", "PUT")

	// router.HandleFunc("/api/Rating/student/sent/{CreatorID}", postedRatings).Methods("GET")

	// router.HandleFunc("/api/Rating/class/sent/{CreatorID}", postedRatings).Methods("GET")

	// router.HandleFunc("/api/Rating/module/sent/{CreatorID}", postedRatings).Methods("GET")

	// router.HandleFunc("/api/Rating/tutor/sent/{CreatorID}", postedRatings).Methods("GET")

	// router.HandleFunc("/api/Rating/received/{CreatorID}", receivedRatings).Methods("GET")

	fmt.Println("Listening at port 5002")
	log.Fatal(http.ListenAndServe(":5002", handlers.CORS(headers, origins, methods)(router)))
}
