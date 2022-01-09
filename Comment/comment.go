package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Comment struct {
	CommentID         int
	StudentID         int
	CommentType       string
	TargetID          int
	CommentData       string
	Anonymous         bool
	DateTimePublished string
}

// type StudentComment struct {
// 	CommentID int
// 	StudentID int

// 	SelectedStudentID int
// 	Comment           string
// 	Anonymous         bool
// 	DateTimePublished string
// }

// type ClassComment struct {
// 	CommentID int
// 	StudentID int

// 	ClassID           int
// 	Comment           string
// 	Anonymous         bool
// 	DateTimePublished string
// }

// type ModuleComment struct {
// 	CommentID int
// 	StudentID int

// 	ModuleID          int
// 	Comment           string
// 	Anonymous         bool
// 	DateTimePublished string
// }

// type TutorComment struct {
// 	CommentID int
// 	StudentID int

// 	ModuleID          int
// 	Comment           string
// 	Anonymous         bool
// 	DateTimePublished string
// }

func getComments(db *sql.DB, commentType string, studentID int) []Comment {
	query := ""
	if commentType == "Student" {
		query = fmt.Sprintf("SELECT * FROM StudentComment WHERE StudentID = '%i'", studentID)
	} else if commentType == "Class" {
		query = fmt.Sprintf("SELECT * FROM ClassComment WHERE StudentID = '%i'", studentID)
	} else if commentType == "Module" {
		query = fmt.Sprintf("SELECT * FROM ModuleComment WHERE StudentID = '%i'", studentID)
	} else if commentType == "Tutor" {
		query = fmt.Sprintf("SELECT * FROM TutorComment WHERE StudentID = '%i'", studentID)
	}
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var commentList []Comment
	for results.Next() {
		var comment Comment
		results.Scan(&comment.CommentID, &comment.StudentID, &comment.TargetID, &comment.CommentData, &comment.Anonymous, comment.DateTimePublished)
		commentList = append(commentList, comment)
	}
	return commentList
}

func postComment(db *sql.DB, comment Comment) {
	StudentID := comment.StudentID
	TargetID := comment.TargetID
	CommentData := comment.CommentData
	Anonymous := comment.Anonymous
	query := ""
	if comment.CommentType == "Student" {
		query = fmt.Sprintf("INSERT INTO StudentComment (StudentID, SelectedStudentID, CommentData, Anonymous) VALUES ('%d', '%d', '%d', '%b')",
			StudentID, TargetID, CommentData, Anonymous)
	} else if comment.CommentType == "Class" {
		query = fmt.Sprintf("INSERT INTO ClassComment (StudentID, ClassID, CommentData, Anonymous) VALUES ('%d', '%d', '%d', '%b')",
			StudentID, TargetID, CommentData, Anonymous)
	} else if comment.CommentType == "Module" {
		query = fmt.Sprintf("INSERT INTO ModuleComment (StudentID, ModuleID, CommentData, Anonymous) VALUES ('%d', '%d', '%d', '%b')",
			StudentID, TargetID, CommentData, Anonymous)
	} else if comment.CommentType == "Tutor" {
		query = fmt.Sprintf("INSERT INTO TutorComment (StudentID, TutorID, CommentData, Anonymous) VALUES ('%d', '%d', '%d', '%b')",
			StudentID, TargetID, CommentData, Anonymous)
	}
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}

//db function to update passenger object details inside
func updateRecord(db *sql.DB, comment Comment) {
	CommentID := comment.CommentID
	CommentData := comment.CommentData
	query := ""
	if comment.CommentType == "Student" {
		query = fmt.Sprintf("UPDATE StudentComment SET CommentData = '%s' WHERE CommentID = '%d'", CommentData, CommentID)
	} else if comment.CommentType == "Class" {
		query = fmt.Sprintf("UPDATE ClassComment SET CommentData = '%s' WHERE CommentID = '%d'", CommentData, CommentID)
	} else if comment.CommentType == "Module" {
		query = fmt.Sprintf("UPDATE ModuleComment SET CommentData = '%s' WHERE CommentID = '%d'", CommentData, CommentID)
	} else if comment.CommentType == "Tutor" {
		query = fmt.Sprintf("UPDATE TutorComment SET CommentData = '%s' WHERE CommentID = '%d'", CommentData, CommentID)
	}
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}
