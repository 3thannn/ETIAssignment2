package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Comment struct {
	CommentID         int
	StudentID         int
	CommentType       string
	TargetID          int
	Comment           string
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

func getComment(db *sql.DB, commentType string) Comment {

}
