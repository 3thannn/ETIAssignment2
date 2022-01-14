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

type Comment struct {
	CommentID         int
	CreatorID         int
	CreatorType       string
	CommentType       string
	TargetID          int
	CommentData       string
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

//Get all comments posted
func getStudentComments(db *sql.DB, CreatorID int) []Comment {
	studentquery := fmt.Sprintf("SELECT * FROM Comment WHERE CreatorID = '%i' AND CommentType = Student", CreatorID)

	studentresults, err := db.Query(studentquery)
	if err != nil {
		panic(err.Error())
	}

	var studentCommentList []Comment
	for studentresults.Next() {
		var comment Comment
		studentresults.Scan(&comment.CommentID, &comment.CreatorID, &comment.TargetID, &comment.CommentType, &comment.CommentData, &comment.Anonymous, comment.DateTimePublished)
		studentCommentList = append(studentCommentList, comment)
	}

	return studentCommentList
}

//STILL NEED TO DO GET STUDENT COMMENTS FROM TUTOR

func getClassComments(db *sql.DB, CreatorID int) []Comment {
	classquery := fmt.Sprintf("SELECT * FROM Comment WHERE CreatorID = '%i' AND CommentType = Class", CreatorID)
	classresults, err := db.Query(classquery)
	if err != nil {
		panic(err.Error())
	}
	var classCommentList []Comment
	for classresults.Next() {
		var comment Comment
		classresults.Scan(&comment.CommentID, &comment.CreatorID, &comment.TargetID, &comment.CommentType, &comment.CommentData, &comment.Anonymous, comment.DateTimePublished)
		classCommentList = append(classCommentList, comment)
	}
	return classCommentList
}

func getModuleComments(db *sql.DB, CreatorID int) []Comment {
	modulequery := fmt.Sprintf("SELECT * FROM Comment WHERE CreatorID = '%i' AND CommentType = Module", CreatorID)
	moduleresults, err := db.Query(modulequery)
	if err != nil {
		panic(err.Error())
	}
	var moduleCommentList []Comment
	for moduleresults.Next() {
		var comment Comment
		moduleresults.Scan(&comment.CommentID, &comment.CreatorID, &comment.TargetID, &comment.CommentType, &comment.CommentData, &comment.Anonymous, comment.DateTimePublished)
		moduleCommentList = append(moduleCommentList, comment)
	}
	return moduleCommentList
}

func getTutorComments(db *sql.DB, CreatorID int) []Comment {
	tutorquery := fmt.Sprintf("SELECT * FROM Comment WHERE CreatorID = '%i' AND CommentType = Tutor", CreatorID)
	tutorresults, err := db.Query(tutorquery)
	if err != nil {
		panic(err.Error())
	}
	var tutorCommentList []Comment
	for tutorresults.Next() {
		var comment Comment
		tutorresults.Scan(&comment.CommentID, &comment.CreatorID, &comment.TargetID, &comment.CommentType, &comment.CommentData, &comment.Anonymous, comment.DateTimePublished)
		tutorCommentList = append(tutorCommentList, comment)
	}
	return tutorCommentList
}

//Get all comments received
func getReceivedComments(db *sql.DB, CreatorID int) []Comment {

	query := fmt.Sprintf("SELECT * FROM StudentComment WHERE TargetCreatorID = '%i'", CreatorID)

	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var commentList []Comment
	for results.Next() {
		var comment Comment
		results.Scan(&comment.CommentID, &comment.CreatorID, &comment.TargetID, &comment.CommentType, &comment.CommentData, &comment.Anonymous, comment.DateTimePublished)
		commentList = append(commentList, comment)
	}
	return commentList
}

func postComment(db *sql.DB, comment Comment) {
	CreatorID := comment.CreatorID
	println(comment.CreatorID)
	TargetID := comment.TargetID
	println(comment.TargetID)
	CommentData := comment.CommentData
	fmt.Printf(comment.CommentData)
	Anonymous := comment.Anonymous
	CommentType := comment.CommentType
	println(comment.Anonymous)
	query := ""
	if comment.CommentType == "Student" {
		query = fmt.Sprintf("INSERT INTO Comment (CreatorID, TargetID, CommentType, CommentData, Anonymous) VALUES ('%d', '%d', '%s', '%s', '%b')",
			CreatorID, TargetID, CommentType, CommentData, Anonymous)
	} else if comment.CommentType == "Class" {
		query = fmt.Sprintf("INSERT INTO Comment (CreatorID, TargetID, CommentType, CommentData, Anonymous) VALUES ('%d', '%d', '%s', '%s', '%b')",
			CreatorID, TargetID, CommentType, CommentData, Anonymous)
	} else if comment.CommentType == "Module" {
		query = fmt.Sprintf("INSERT INTO Comment (CreatorID, TargetID, CommentType, CommentData, Anonymous) VALUES ('%d', '%d', '%s', '%s', '%b')",
			CreatorID, TargetID, CommentType, CommentData, Anonymous)
	} else if comment.CommentType == "Tutor" {
		query = fmt.Sprintf("INSERT INTO Comment (CreatorID, TargetID, CommentType, CommentData, Anonymous) VALUES ('%d', '%d', '%s', '%s', '%b')",
			CreatorID, TargetID, CommentType, CommentData, Anonymous)
	}
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}

func updateRecord(db *sql.DB, comment Comment) {
	CommentID := comment.CommentID
	CommentData := comment.CommentData
	query := ""
	if comment.CommentType == "Student" {
		query = fmt.Sprintf("UPDATE Comment SET CommentData = '%s' WHERE CommentID = '%d' AND CommentType = Student", CommentData, CommentID)
	} else if comment.CommentType == "Class" {
		query = fmt.Sprintf("UPDATE  Comment SET CommentData = '%s' WHERE CommentID = '%d' AND CommentType = Class", CommentData, CommentID)
	} else if comment.CommentType == "Module" {
		query = fmt.Sprintf("UPDATE Comment SET CommentData = '%s' WHERE CommentID = '%d' AND CommentType = Module", CommentData, CommentID)
	} else if comment.CommentType == "Tutor" {
		query = fmt.Sprintf("UPDATE Comment SET CommentData = '%s' WHERE CommentID = '%d' AND CommentType = Tutor", CommentData, CommentID)
	}
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}

func comment(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "POST" {
		var newComment Comment
		reqBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(reqBody, &newComment)
			postComment(db, newComment)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("201 - Comment Posted!"))
		}
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("422 - Please enter trip detais in JSON format!"))
	}
}

//Get all comments received
func receivedComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment")
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
			var personalComments []Comment = getReceivedComments(db, CreatorIDint)
			if len(personalComments) > 0 {
				fmt.Println(personalComments)
				json.NewEncoder(w).Encode(personalComments)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
	}
	if r.Method == "PUT" {
		// read the string sent to the service
		var commentInput Comment
		reqBody, err := ioutil.ReadAll(r.Body)

		if err == nil {
			// convert JSON to object
			json.Unmarshal(reqBody, &commentInput)
			updateRecord(db, commentInput)
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("202 - Your comment has been successfully updated!"))
		}
	} else {
		w.WriteHeader(
			http.StatusUnprocessableEntity)
		w.Write([]byte("422 - Please enter account detais in JSON format!"))
	}
}

func postedStudentComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment")
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
			studentCommentList := getStudentComments(db, CreatorIDint)
			if len(studentCommentList) > 0 {
				fmt.Println(studentCommentList)
				json.NewEncoder(w).Encode(studentCommentList)
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

	router.HandleFunc("/api/comment", comment).Methods("POST", "PUT")

	// router.HandleFunc("/api/comment/student/sent/{CreatorID}", postedComments).Methods("GET")

	// router.HandleFunc("/api/comment/class/sent/{CreatorID}", postedComments).Methods("GET")

	// router.HandleFunc("/api/comment/module/sent/{CreatorID}", postedComments).Methods("GET")

	// router.HandleFunc("/api/comment/tutor/sent/{CreatorID}", postedComments).Methods("GET")

	// router.HandleFunc("/api/comment/received/{CreatorID}", receivedComments).Methods("GET")

	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", handlers.CORS(headers, origins, methods)(router)))
}
