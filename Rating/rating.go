package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Rating struct {
	RatingID          int
	StudentID         int
	RatingType        string
	TargetID          int
	RatingScore       string
	Anonymous         bool
	DateTimePublished string
}

// type StudentRating struct {
// 	CommentID int
// 	StudentID int

// 	SelectedStudentID int
// 	Rating            string
// 	Anonymous         bool
// 	DateTimePublished string
// }

// type ClassRating struct {
// 	CommentID int
// 	StudentID int

// 	ClassID           int
// 	Rating            string
// 	Anonymous         bool
// 	DateTimePublished string
// }

// type ModuleRating struct {
// 	CommentID int
// 	StudentID int

// 	ModuleID          int
// 	Rating            string
// 	Anonymous         bool
// 	DateTimePublished string
// }

// type TutorRating struct {
// 	CommentID int
// 	StudentID int

// 	TutorID           int
// 	Rating            string
// 	Anonymous         bool
// 	DateTimePublished string
// }

func getComments(db *sql.DB, ratingType string, studentID int) []Rating {
	query := ""
	if ratingType == "Student" {
		query = fmt.Sprintf("SELECT * FROM StudenrRating WHERE StudentID = '%i'", studentID)
	} else if ratingType == "Class" {
		query = fmt.Sprintf("SELECT * FROM ClassRating WHERE StudentID = '%i'", studentID)
	} else if ratingType == "Module" {
		query = fmt.Sprintf("SELECT * FROM ModuleRating WHERE StudentID = '%i'", studentID)
	} else if ratingType == "Tutor" {
		query = fmt.Sprintf("SELECT * FROM TutorRating WHERE StudentID = '%i'", studentID)
	}
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var ratingList []Rating
	for results.Next() {
		var rating Rating
		results.Scan(&rating.RatingID, &rating.StudentID, &rating.TargetID, &rating.RatingScore, &rating.Anonymous, rating.DateTimePublished)
		ratingList = append(ratingList, rating)
	}
	return ratingList
}

func postComment(db *sql.DB, rating Rating) {
	StudentID := rating.StudentID
	TargetID := rating.TargetID
	RatingScore := rating.RatingScore
	Anonymous := rating.Anonymous
	query := ""
	if rating.RatingType == "Student" {
		query = fmt.Sprintf("INSERT INTO StudentRating (StudentID, SelectedStudentID, RatingScore, Anonymous) VALUES ('%d', '%d', '%d', '%b')",
			StudentID, TargetID, RatingScore, Anonymous)
	} else if rating.RatingType == "Class" {
		query = fmt.Sprintf("INSERT INTO ClassRating (StudentID, ClassID, RatingScore, Anonymous) VALUES ('%d', '%d', '%d', '%b')",
			StudentID, TargetID, RatingScore, Anonymous)
	} else if rating.RatingType == "Module" {
		query = fmt.Sprintf("INSERT INTO ModuleRating (StudentID, ModuleID, RatingScore, Anonymous) VALUES ('%d', '%d', '%d', '%b')",
			StudentID, TargetID, RatingScore, Anonymous)
	} else if rating.RatingType == "Tutor" {
		query = fmt.Sprintf("INSERT INTO TutorRating (StudentID, TutorID, RatingScore, Anonymous) VALUES ('%d', '%d', '%d', '%b')",
			StudentID, TargetID, RatingScore, Anonymous)
	}
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}

//db function to update passenger object details inside
func updateRecord(db *sql.DB, rating Rating) {
	RatingID := rating.RatingID
	RatingScore := rating.RatingScore
	query := ""
	if rating.RatingType == "Student" {
		query = fmt.Sprintf("UPDATE StudentRating SET RatingScore = '%d' WHERE CommentID = '%d'", RatingScore, RatingID)
	} else if rating.RatingType == "Class" {
		query = fmt.Sprintf("UPDATE ClassRating SET RatingScore = '%d' WHERE CommentID = '%d'", RatingScore, RatingID)
	} else if rating.RatingType == "Module" {
		query = fmt.Sprintf("UPDATE ModuleRating SET RatingScore = '%d' WHERE CommentID = '%d'", RatingScore, RatingID)
	} else if rating.RatingType == "Tutor" {
		query = fmt.Sprintf("UPDATE TutorRating SET RatingScore = '%d' WHERE CommentID = '%d'", RatingScore, RatingID)
	}
	_, err := db.Query(query) //Run Query

	if err != nil {
		panic(err.Error())
	}
}
