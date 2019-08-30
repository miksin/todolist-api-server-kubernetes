package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

var PORT = 8080
var DB_NAME = "go_todo"
var DB_USER = "root"
var DB_PASSWORD = "password"
var DB_HOST = "localhost"
var DB_PORT = 3306

func initEnv() {
	if v := os.Getenv("PORT"); len(v) > 0 {
		if p, ok := strconv.Atoi(v); ok == nil {
			PORT = p
		}
	}
	if v := os.Getenv("DB_NAME"); len(v) > 0 {
		DB_NAME = v
	}
	if v := os.Getenv("DB_NAME"); len(v) > 0 {
		DB_NAME = v
	}
	if v := os.Getenv("DB_NAME"); len(v) > 0 {
		DB_NAME = v
	}
	if v := os.Getenv("DB_HOST"); len(v) > 0 {
		DB_HOST = v
	}
	if v := os.Getenv("DB_PORT"); len(v) > 0 {
		if p, ok := strconv.Atoi(v); ok == nil {
			DB_PORT = p
		}
	}
}

func main() {
	initEnv()
	InitRepo()
	defer CloseRepo()

	router := NewRouter()

	log.Printf("server listen on: %s\n", strconv.Itoa(PORT))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), router))
}
