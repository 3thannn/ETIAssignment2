package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type Module struct {
	ModuleID int
	Name     string
}
