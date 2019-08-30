package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func responseJson(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1>", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos, err := RepoListTodo()
	if err != nil {
		responseJson(w, http.StatusInternalServerError, Msg{Msg: err.Error()})
		return
	}
	responseJson(w, http.StatusOK, todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, ok := strconv.Atoi(vars["todoId"])
	if ok != nil {
		responseJson(w, http.StatusInternalServerError, Msg{Msg: "server error"})
		return
	}

	if t, err := RepoFindTodo(todoId); err == nil {
		responseJson(w, http.StatusOK, t)
	} else {
		responseJson(w, http.StatusNotFound, Msg{Msg: err.Error()})
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024*1024))
	if err != nil {
		responseJson(w, http.StatusInternalServerError, Msg{Msg: "server error"})
		return
	}
	if err := r.Body.Close(); err != nil {
		responseJson(w, http.StatusInternalServerError, Msg{Msg: "server error"})
		return
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		responseJson(w, http.StatusUnprocessableEntity, Msg{Msg: "body error"})
		return
	}

	if t, err := RepoCreateTodo(todo); err != nil {
		responseJson(w, http.StatusInternalServerError, Msg{Msg: err.Error()})
	} else {
		responseJson(w, http.StatusCreated, t)
	}
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, ok := strconv.Atoi(vars["todoId"])
	if ok != nil {
		responseJson(w, http.StatusInternalServerError, Msg{Msg: "server error"})
		return
	}

	if err := RepoDestroyTodo(todoId); err == nil {
		responseJson(w, http.StatusOK, Msg{Msg: "success"})
	} else {
		responseJson(w, http.StatusNotFound, Msg{Msg: err.Error()})
	}
}
