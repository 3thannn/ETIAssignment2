package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type StudentRating struct {
	CommentID int
	StudentID int

	SelectedStudentID int
	Rating            string
	Anonymous         bool
	DateTimePublished string
}

type ClassRating struct {
	CommentID int
	StudentID int

	ClassID           int
	Rating            string
	Anonymous         bool
	DateTimePublished string
}

type ModuleRating struct {
	CommentID int
	StudentID int

	ModuleID          int
	Rating            string
	Anonymous         bool
	DateTimePublished string
}

type TutorRating struct {
	CommentID int
	StudentID int

	TutorID           int
	Rating            string
	Anonymous         bool
	DateTimePublished string
}
