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
	TargetID          int
	TargetType        string
	CommentData       string
	Anonymous         int
	DateTimePublished string
	CreatorName       string
	TargetName        string
}

type Object struct {
	ID   int
	Name string
}

//Gets all Student's ID which is tied to StudentID
func getAllStudents(db *sql.DB) []Object {
	url := "http://localhost:5003/api/student"
	response, err := http.Get(url)
	var studentList []Object
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode == http.StatusNotFound {
			fmt.Println("409 - No Students Found!")
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			response.Body.Close()
			json.Unmarshal(data, &studentList)
			fmt.Println("202 - Successfully received Students!")
		}
	}
	return studentList
}

//Gets all Tutor's Names which is tied to TutorID
func getAllTutors(db *sql.DB) []Object {
	url := "http://localhost:5004/api/tutor"
	response, err := http.Get(url)
	var tutorList []Object
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode == http.StatusNotFound {
			fmt.Println("409 - No Tutors Found!")
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			response.Body.Close()
			json.Unmarshal(data, &tutorList)
			fmt.Println("202 - Successfully received Tutors!")
		}
	}
	return tutorList
}

//Gets all Class Name which is tied to ClassID
func getAllClasses(db *sql.DB) []Object {
	url := "http://localhost:5005/api/class"
	response, err := http.Get(url)
	var classList []Object
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode == http.StatusNotFound {
			fmt.Println("409 - No Classes Found!")
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			response.Body.Close()
			json.Unmarshal(data, &classList)
			fmt.Println("202 - Successfully received Classes!")
		}
	}
	return classList
}

//Gets all Module Name which is tied to ModuleID
func getAllModules(db *sql.DB) []Object {
	url := "http://localhost:5006/api/module"
	response, err := http.Get(url)
	var moduleList []Object
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode == http.StatusNotFound {
			fmt.Println("409 - No Modules Found!")
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			response.Body.Close()
			json.Unmarshal(data, &moduleList)
			fmt.Println("202 - Successfully received Modules!")
		}
	}
	return moduleList
}

func linkStudentToID(id int, studentList []Object) Object {
	var student Object
	for _, student := range studentList {
		if student.ID == id {
			return student
		}
	}
	return student
}

func linkTutortToID(id int, tutorList []Object) Object {
	var tutor Object
	for _, tutor := range tutorList {
		if tutor.ID == id {
			return tutor
		}
	}
	return tutor
}

func linkClasstToID(id int, classList []Object) Object {
	var class Object
	for _, class := range classList {
		if class.ID == id {
			return class
		}
	}
	return class
}

func linkModuletToID(id int, moduleList []Object) Object {
	var module Object
	for _, module := range moduleList {
		if module.ID == id {
			return module
		}
	}
	return module
}

//3.9.1 View comments
//Get all comments to students
func getStudentComments(db *sql.DB) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	studentquery := "SELECT * FROM Comment WHERE TargetType = 'Student';"

	studentresults, err := db.Query(studentquery)
	if err != nil {
		panic(err.Error())
	}
	var studentCommentList []Comment
	for studentresults.Next() {
		var comment Comment
		studentresults.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)
		if comment.Anonymous == 0 {
			if comment.CreatorType == "Student" {
				println("This ran here")
				student := linkStudentToID(comment.CreatorID, studentList)
				comment.CreatorName = student.Name
			} else if comment.CreatorType == "Tutor" {
				tutor := linkTutortToID(comment.CreatorID, tutorList)
				comment.CreatorName = tutor.Name
			}
		}
		student := linkStudentToID(comment.TargetID, studentList)
		comment.TargetName = student.Name
		fmt.Println(comment)
		studentCommentList = append(studentCommentList, comment)
	}

	return studentCommentList
}

//3.9.1 View comments
//Get all comments to classes
func getClassComments(db *sql.DB) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	classList := getAllClasses(db)
	classQuery := "SELECT * FROM Comment WHERE TargetType = 'Class';"

	classResults, err := db.Query(classQuery)
	if err != nil {
		panic(err.Error())
	}
	var classCommentList []Comment
	for classResults.Next() {
		var comment Comment
		classResults.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)
		if comment.Anonymous == 0 {
			if comment.CreatorType == "Student" {
				student := linkStudentToID(comment.CreatorID, studentList)
				fmt.Println(student)
				comment.CreatorName = student.Name
			} else if comment.CreatorType == "Tutor" {
				tutor := linkTutortToID(comment.CreatorID, tutorList)
				comment.CreatorName = tutor.Name
			}
		}
		class := linkClasstToID(comment.TargetID, classList)
		comment.TargetName = class.Name
		classCommentList = append(classCommentList, comment)

	}

	return classCommentList
}

//3.9.1 View comments
//Get all comments to modules
func getModuleComments(db *sql.DB) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	moduleList := getAllModules(db)
	moduleQuery := "SELECT * FROM Comment WHERE TargetType = 'Module';"

	moduleResults, err := db.Query(moduleQuery)
	if err != nil {
		panic(err.Error())
	}
	var moduleCommentList []Comment
	for moduleResults.Next() {
		var comment Comment
		moduleResults.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)
		if comment.Anonymous == 0 {
			if comment.CreatorType == "Student" {
				student := linkStudentToID(comment.CreatorID, studentList)
				comment.CreatorName = student.Name
			} else if comment.CreatorType == "Tutor" {
				tutor := linkTutortToID(comment.CreatorID, tutorList)
				comment.CreatorName = tutor.Name
			}
		}
		module := linkModuletToID(comment.TargetID, moduleList)
		comment.TargetName = module.Name
		fmt.Println(comment)
		moduleCommentList = append(moduleCommentList, comment)

	}

	return moduleCommentList
}

//3.9.1 View comments
//Get all comments to Tutors
func getTutorComments(db *sql.DB) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	tutorQuery := "SELECT * FROM Comment WHERE TargetType = 'Tutor';"

	tutorResults, err := db.Query(tutorQuery)
	if err != nil {
		panic(err.Error())
	}
	var tutorCommentList []Comment
	for tutorResults.Next() {
		var comment Comment
		tutorResults.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)
		if comment.Anonymous == 0 {
			if comment.CreatorType == "Student" {
				student := linkStudentToID(comment.CreatorID, studentList)
				comment.CreatorName = student.Name
			} else if comment.CreatorType == "Tutor" {
				tutor := linkTutortToID(comment.CreatorID, tutorList)
				comment.CreatorName = tutor.Name
			}
		}
		tutor := linkTutortToID(comment.TargetID, tutorList)
		comment.TargetName = tutor.Name
		fmt.Println(comment)
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
		results.Scan(&comment.CommentID, &comment.CreatorID, &comment.TargetID, &comment.TargetType, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)
		commentList = append(commentList, comment)
	}
	return commentList
}

func postComment(db *sql.DB, comment Comment) {
	CreatorType := comment.CreatorType
	println(comment.CreatorType)
	CreatorID := comment.CreatorID
	println(comment.CreatorID)
	TargetID := comment.TargetID
	println(comment.TargetID)
	CommentData := comment.CommentData
	println(comment.CommentData)
	Anonymous := comment.Anonymous
	println(comment.Anonymous)
	TargetType := comment.TargetType
	println(comment.TargetType)
	query := fmt.Sprintf("INSERT INTO Comment (CreatorType, CreatorID, TargetID, TargetType, CommentData, Anonymous, DateTimePublished) VALUES ('%s', '%d', '%d', '%s', '%s', '%b', NOW())",
		CreatorType, CreatorID, TargetID, TargetType, CommentData, Anonymous)
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}

func updateRecord(db *sql.DB, comment Comment) {
	CommentID := comment.CommentID
	CommentData := comment.CommentData
	query := ""
	if comment.TargetType == "Student" {
		query = fmt.Sprintf("UPDATE Comment SET CommentData = '%s' WHERE CommentID = '%d' AND TargetType = Student", CommentData, CommentID)
	} else if comment.TargetType == "Class" {
		query = fmt.Sprintf("UPDATE  Comment SET CommentData = '%s' WHERE CommentID = '%d' AND TargetType = Class", CommentData, CommentID)
	} else if comment.TargetType == "Module" {
		query = fmt.Sprintf("UPDATE Comment SET CommentData = '%s' WHERE CommentID = '%d' AND TargetType = Module", CommentData, CommentID)
	} else if comment.TargetType == "Tutor" {
		query = fmt.Sprintf("UPDATE Comment SET CommentData = '%s' WHERE CommentID = '%d' AND TargetType = Tutor", CommentData, CommentID)
	}
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}

func comment(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment?parseTime=true")
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
		w.Write([]byte("422 - Comment Info not in JSON format!"))
	}
}
func studentComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment?parseTime=true")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			studentCommentList := getStudentComments(db)
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
func tutorComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment?parseTime=true")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			tutorCommentList := getTutorComments(db)
			if len(tutorCommentList) > 0 {
				fmt.Println(tutorCommentList)
				json.NewEncoder(w).Encode(tutorCommentList)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
	}
}
func classComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment?parseTime=true")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			classCommentList := getClassComments(db)
			if len(classCommentList) > 0 {
				fmt.Println(classCommentList)
				json.NewEncoder(w).Encode(classCommentList)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
	}
}
func moduleComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment?parseTime=true")
	// handle error
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			moduleCommentList := getModuleComments(db)
			if len(moduleCommentList) > 0 {
				fmt.Println(moduleCommentList)
				json.NewEncoder(w).Encode(moduleCommentList)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		} else {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}
	}
}

//Get all comments received
func receivedComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment?parseTime=true")
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
	// db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIAssignment2Comment")
	// // handle error
	// if err != nil {
	// 	panic(err.Error())
	// }

	// if err != nil {
	// 	panic(err.Error())
	// }

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

	router.HandleFunc("/api/comment/student", studentComments).Methods("GET")

	router.HandleFunc("/api/comment/tutor", tutorComments).Methods("GET")

	router.HandleFunc("/api/comment/class", classComments).Methods("GET")

	router.HandleFunc("/api/comment/module", moduleComments).Methods("GET")

	// router.HandleFunc("/api/comment/student/sent/{CreatorID}", postedComments).Methods("GET")

	// router.HandleFunc("/api/comment/class/sent/{CreatorID}", postedComments).Methods("GET")

	// router.HandleFunc("/api/comment/module/sent/{CreatorID}", postedComments).Methods("GET")

	// router.HandleFunc("/api/comment/tutor/sent/{CreatorID}", postedComments).Methods("GET")

	// router.HandleFunc("/api/comment/received/{CreatorID}", receivedComments).Methods("GET")

	fmt.Println("Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", handlers.CORS(headers, origins, methods)(router)))
}
