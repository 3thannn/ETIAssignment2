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
	url := "http://studentcontainer:9043/api/student"
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
	url := "http://tutorcontainer:9044/api/tutor"
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
	url := "http://classcontainer:9046/api/class"
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
	url := "http://modulecontainer:9045/api/module"
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
func linkStudentToID(db *sql.DB, id int, studentList []Object) Object {
	var student Object
	for _, student := range studentList {
		if student.ID == id {
			return student
		}
	}
	return student
}

func linkTutorToID(db *sql.DB, id int, tutorList []Object) Object {
	var tutor Object
	for _, tutor := range tutorList {
		if tutor.ID == id {
			return tutor
		}
	}
	return tutor
}

func linkClassToID(db *sql.DB, id int, classList []Object) Object {
	var class Object
	for _, class := range classList {
		if class.ID == id {
			return class
		}
	}
	return class
}

func linkModuleToID(db *sql.DB, id int, moduleList []Object) Object {
	var module Object
	for _, module := range moduleList {
		if module.ID == id {
			return module
		}
	}
	return module
}

func getComment(db *sql.DB, CommentId int) Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	moduleList := getAllModules(db)
	classList := getAllClasses(db)
	query := fmt.Sprintf("SELECT * FROM Comment WHERE CommentID = '%d';", CommentId)
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var comment Comment
	for results.Next() {

		results.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)
		if comment.CreatorType == "Student" {
			student := linkStudentToID(db, comment.CreatorID, studentList)
			comment.CreatorName = student.Name
			println(student.Name)
		} else if comment.CreatorType == "Tutor" {
			tutor := linkTutorToID(db, comment.CreatorID, tutorList)
			comment.CreatorName = tutor.Name
			println(tutor.Name)
		}
		if comment.TargetType == "Student" {
			student := linkStudentToID(db, comment.TargetID, studentList)
			comment.TargetName = student.Name
		} else if comment.TargetType == "Tutor" {
			tutor := linkTutorToID(db, comment.TargetID, tutorList)
			comment.TargetName = tutor.Name
		} else if comment.TargetType == "Module" {
			module := linkModuleToID(db, comment.TargetID, moduleList)
			fmt.Print("this is the module: ")
			fmt.Println(module)
			comment.TargetName = module.Name
		} else if comment.TargetType == "Class" {
			class := linkClassToID(db, comment.TargetID, classList)
			comment.TargetName = class.Name
		}
		fmt.Println(comment)
	}
	return comment
}

//3.9.1 View comments
//Get all comments to students
func getStudentComments(db *sql.DB, targetID int) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	fmt.Println(studentList)
	fmt.Println(tutorList)
	studentquery := fmt.Sprintf("SELECT * FROM Comment WHERE TargetType = 'Student' AND TargetID = '%d';", targetID)

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
				student := linkStudentToID(db, comment.CreatorID, studentList)
				comment.CreatorName = student.Name
				println(student.Name)
			} else if comment.CreatorType == "Tutor" {
				tutor := linkTutorToID(db, comment.CreatorID, tutorList)
				comment.CreatorName = tutor.Name
				println(tutor.Name)
			}
		}
		student := linkStudentToID(db, comment.TargetID, studentList)
		comment.TargetName = student.Name
		fmt.Println(comment)
		studentCommentList = append(studentCommentList, comment)
	}

	return studentCommentList
}

//3.9.1 View comments
//Get all comments to classes
func getClassComments(db *sql.DB, targetID int) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	classList := getAllClasses(db)
	classQuery := fmt.Sprintf("SELECT * FROM Comment WHERE TargetType = 'Class' AND TargetID = '%d'", targetID)

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
				student := linkStudentToID(db, comment.CreatorID, studentList)
				fmt.Println(student)
				comment.CreatorName = student.Name
			} else if comment.CreatorType == "Tutor" {
				tutor := linkTutorToID(db, comment.CreatorID, tutorList)
				comment.CreatorName = tutor.Name
			}
		}
		class := linkClassToID(db, comment.TargetID, classList)
		comment.TargetName = class.Name
		classCommentList = append(classCommentList, comment)

	}

	return classCommentList
}

//3.9.1 View comments
//Get all comments to modules
func getModuleComments(db *sql.DB, targetID int) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	moduleList := getAllModules(db)
	moduleQuery := fmt.Sprintf("SELECT * FROM Comment WHERE TargetType = 'Module'; AND TargetID = '%d'", targetID)

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
				student := linkStudentToID(db, comment.CreatorID, studentList)
				comment.CreatorName = student.Name
			} else if comment.CreatorType == "Tutor" {
				tutor := linkTutorToID(db, comment.CreatorID, tutorList)
				comment.CreatorName = tutor.Name
			}
		}
		module := linkModuleToID(db, comment.TargetID, moduleList)
		comment.TargetName = module.Name
		fmt.Println(comment)
		moduleCommentList = append(moduleCommentList, comment)

	}

	return moduleCommentList
}

//3.9.1 View comments
//Get all comments to Tutors
func getTutorComments(db *sql.DB, targetID int) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	fmt.Println(studentList)
	fmt.Println(tutorList)
	tutorquery := fmt.Sprintf("SELECT * FROM Comment WHERE TargetType = 'Tutor' AND TargetID = '%d';", targetID)

	tutorResults, err := db.Query(tutorquery)
	if err != nil {
		panic(err.Error())
	}
	var tutorCommentList []Comment
	for tutorResults.Next() {
		var comment Comment
		tutorResults.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)
		fmt.Print(comment)
		if comment.Anonymous == 0 {
			if comment.CreatorType == "Student" {
				student := linkStudentToID(db, comment.CreatorID, studentList)
				comment.CreatorName = student.Name
				println(student.Name)
			} else if comment.CreatorType == "Tutor" {
				tutor := linkTutorToID(db, comment.CreatorID, tutorList)
				comment.CreatorName = tutor.Name
				println(tutor.Name)
			}
		}
		tutor := linkTutorToID(db, comment.TargetID, tutorList)
		comment.TargetName = tutor.Name
		fmt.Println(comment)
		tutorCommentList = append(tutorCommentList, comment)
	}

	return tutorCommentList
}

func insertComment(db *sql.DB, comment Comment) {
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
	fmt.Println(comment)
	CommentID := comment.CommentID
	CommentData := comment.CommentData
	Anonymous := comment.Anonymous
	query := fmt.Sprintf("UPDATE Comment SET CommentData = '%s', Anonymous = '%d' WHERE CommentID = '%d' ", CommentData, Anonymous, CommentID)
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}

//Get all comments received
func getReceivedComments(db *sql.DB, TargetType string, TargetID int) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	moduleList := getAllModules(db)
	classList := getAllClasses(db)
	query := fmt.Sprintf("SELECT * FROM Comment WHERE TargetID = '%d' AND TargetType = '%s'", TargetID, TargetType)

	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var commentList []Comment
	for results.Next() {
		var comment Comment
		results.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)
		if comment.CreatorType == "Student" {
			student := linkStudentToID(db, comment.CreatorID, studentList)
			comment.CreatorName = student.Name
			println(student.Name)
		} else if comment.CreatorType == "Tutor" {
			tutor := linkTutorToID(db, comment.CreatorID, tutorList)
			comment.CreatorName = tutor.Name
			println(tutor.Name)
		}
		if comment.TargetType == "Student" {
			student := linkStudentToID(db, comment.TargetID, studentList)
			comment.TargetName = student.Name
		} else if comment.TargetType == "Tutor" {
			tutor := linkTutorToID(db, comment.TargetID, tutorList)
			comment.TargetName = tutor.Name
		} else if comment.TargetType == "Module" {
			module := linkModuleToID(db, comment.TargetID, moduleList)
			fmt.Print("this is the module: ")
			fmt.Println(module)
			comment.TargetName = module.Name
		} else if comment.TargetType == "Class" {
			class := linkClassToID(db, comment.TargetID, classList)
			comment.TargetName = class.Name
		}
		fmt.Println(comment)
		commentList = append(commentList, comment)
	}
	return commentList
}

func getPostedComments(db *sql.DB, CreatorType string, CreatorID int) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	moduleList := getAllModules(db)
	classList := getAllClasses(db)
	query := fmt.Sprintf("SELECT * FROM Comment WHERE CreatorID = '%d' AND CreatorType = '%s'", CreatorID, CreatorType)
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var commentList []Comment
	for results.Next() {
		var comment Comment
		results.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)
		if comment.CreatorType == "Student" {
			student := linkStudentToID(db, comment.CreatorID, studentList)
			comment.CreatorName = student.Name
			println(student.Name)
		} else if comment.CreatorType == "Tutor" {
			tutor := linkTutorToID(db, comment.CreatorID, tutorList)
			comment.CreatorName = tutor.Name
			println(tutor.Name)
		}
		if comment.TargetType == "Student" {
			student := linkStudentToID(db, comment.TargetID, studentList)
			comment.TargetName = student.Name
		} else if comment.TargetType == "Tutor" {
			tutor := linkTutorToID(db, comment.TargetID, tutorList)
			comment.TargetName = tutor.Name
		} else if comment.TargetType == "Module" {
			module := linkModuleToID(db, comment.TargetID, moduleList)
			fmt.Print("this is the module: ")
			fmt.Println(module)
			comment.TargetName = module.Name
		} else if comment.TargetType == "Class" {
			class := linkClassToID(db, comment.TargetID, classList)
			comment.TargetName = class.Name
		}
		fmt.Println(comment)
		commentList = append(commentList, comment)
	}
	return commentList
}

//Get all anonymous comments received
func getReceivedAnonymousComments(db *sql.DB, TargetType string, TargetID int) []Comment {
	studentList := getAllStudents(db)
	tutorList := getAllTutors(db)
	moduleList := getAllModules(db)
	classList := getAllClasses(db)
	query := fmt.Sprintf("SELECT * FROM Comment WHERE TargetID = '%d' AND TargetType = '%s' AND Anonymous = True", TargetID, TargetType)

	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var commentList []Comment
	for results.Next() {
		var comment Comment
		results.Scan(&comment.CommentID, &comment.CreatorType, &comment.CreatorID, &comment.TargetType, &comment.TargetID, &comment.CommentData, &comment.Anonymous, &comment.DateTimePublished)
		if comment.CreatorType == "Student" {
			student := linkStudentToID(db, comment.CreatorID, studentList)
			comment.CreatorName = student.Name
			println(student.Name)
		} else if comment.CreatorType == "Tutor" {
			tutor := linkTutorToID(db, comment.CreatorID, tutorList)
			comment.CreatorName = tutor.Name
			println(tutor.Name)
		}
		if comment.TargetType == "Student" {
			student := linkStudentToID(db, comment.TargetID, studentList)
			comment.TargetName = student.Name
		} else if comment.TargetType == "Tutor" {
			tutor := linkTutorToID(db, comment.TargetID, tutorList)
			comment.TargetName = tutor.Name
		} else if comment.TargetType == "Module" {
			module := linkModuleToID(db, comment.TargetID, moduleList)
			fmt.Print("this is the module: ")
			fmt.Println(module)
			comment.TargetName = module.Name
		} else if comment.TargetType == "Class" {
			class := linkClassToID(db, comment.TargetID, classList)
			comment.TargetName = class.Name
		}
		fmt.Println(comment)
		commentList = append(commentList, comment)
	}
	return commentList
}

func comment(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(db1:9047)/ETIAssignment2Comment")
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	TargetID := params["id"]
	TargetIDInt, err := strconv.Atoi(TargetID)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		var comment = getComment(db, TargetIDInt)
		fmt.Println(comment)
		json.NewEncoder(w).Encode(comment)
	}
	if r.Method == "PUT" {
		var comment Comment
		reqBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(reqBody, &comment)
			updateRecord(db, comment)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("201 - Comment Updated!"))
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("422 - Comment Info not in JSON format!"))
		}
	}
}

func postComment(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(db1:9047)/ETIAssignment2Comment")
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "POST" {
		var newComment Comment
		reqBody, err := ioutil.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(reqBody, &newComment)
			insertComment(db, newComment)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("201 - Comment Posted!"))
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("422 - Comment Info not in JSON format!"))
		}
	}
}

func studentComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(db1:9047)/ETIAssignment2Comment")
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	studentID := params["studentid"]
	studentIDint, err := strconv.Atoi(studentID)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		studentCommentList := getStudentComments(db, studentIDint)
		if len(studentCommentList) > 0 {
			fmt.Println(studentCommentList)
			json.NewEncoder(w).Encode(studentCommentList)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
func tutorComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(db1:9047)/ETIAssignment2Comment")
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	tutorID := params["tutorid"]
	tutorIDint, err := strconv.Atoi(tutorID)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		tutorCommentList := getTutorComments(db, tutorIDint)
		if len(tutorCommentList) > 0 {
			fmt.Println(tutorCommentList)
			json.NewEncoder(w).Encode(tutorCommentList)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func classComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(db1:9047)/ETIAssignment2Comment")
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	classID := params["classid"]
	classIDint, err := strconv.Atoi(classID)
	if err != nil {
		panic(err.Error())
	}
	// handle error
	if r.Method == "GET" {
		classCommentList := getClassComments(db, classIDint)
		if len(classCommentList) > 0 {
			fmt.Println(classCommentList)
			json.NewEncoder(w).Encode(classCommentList)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
func moduleComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(db1:9047)/ETIAssignment2Comment")
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
		moduleCommentList := getModuleComments(db, moduleIDint)
		if len(moduleCommentList) > 0 {
			fmt.Println(moduleCommentList)
			json.NewEncoder(w).Encode(moduleCommentList)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

//Get all comments received
func receivedAnonymousComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(db1:9047)/ETIAssignment2Comment")
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	TargetType := params["type"]
	TargetID := params["id"]
	TargetIDInt, err := strconv.Atoi(TargetID)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			var anonymousComments []Comment = getReceivedAnonymousComments(db, TargetType, TargetIDInt)
			if len(anonymousComments) > 0 {
				fmt.Println(anonymousComments)
				json.NewEncoder(w).Encode(anonymousComments)
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
	db, err := sql.Open("mysql", "root:password@tcp(db1:9047)/ETIAssignment2Comment")
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	TargetType := params["type"]
	TargetID := params["id"]
	TargetIDInt, err := strconv.Atoi(TargetID)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			var personalComments []Comment = getReceivedComments(db, TargetType, TargetIDInt)
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
}

//Get all comments received
func postedComments(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(db1:9047)/ETIAssignment2Comment")
	if err != nil {
		panic(err.Error())
	}
	params := mux.Vars(r)
	TargetType := params["type"]
	TargetID := params["id"]
	TargetIDInt, err := strconv.Atoi(TargetID)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		if err == nil {
			var postedComments []Comment = getPostedComments(db, TargetType, TargetIDInt)
			if len(postedComments) > 0 {
				fmt.Println(postedComments)
				json.NewEncoder(w).Encode(postedComments)
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

	router.HandleFunc("/api/comment/{id}", comment).Methods("GET", "PUT")

	router.HandleFunc("/api/comment/create", postComment).Methods("POST")

	router.HandleFunc("/api/comment/received/{type}/{id}", receivedComments).Methods("GET")

	router.HandleFunc("/api/comment/received/anonymous/{type}/{id}", receivedAnonymousComments).Methods("GET")

	router.HandleFunc("/api/comment/posted/{type}/{id}", postedComments).Methods("GET")

	router.HandleFunc("/api/comment/student/{studentid}", studentComments).Methods("GET")

	router.HandleFunc("/api/comment/tutor/{tutorid}", tutorComments).Methods("GET")

	router.HandleFunc("/api/comment/class/{classid}", classComments).Methods("GET")

	router.HandleFunc("/api/comment/module/{moduleid}", moduleComments).Methods("GET")

	fmt.Println("Listening at port 9041")
	log.Fatal(http.ListenAndServe(":9041", handlers.CORS(headers, origins, methods)(router)))
}
