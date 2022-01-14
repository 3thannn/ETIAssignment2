package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type Class struct {
	ClassID int
	Name    string
}
