package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type Tutor struct {
	TutorID int
	Name    string
}
