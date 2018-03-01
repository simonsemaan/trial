package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/simonsemaan/trial/Cassandra"
	"github.com/simonsemaan/trial/User_information"
	"github.com/simonsemaan/trial/Users"
)

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func main() {
	CassandraSession := Cassandra.Session
	defer CassandraSession.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartbeat)
	router.HandleFunc("/users/new", Users.Post)
	router.HandleFunc("/users", Users.Get)
	router.HandleFunc("/user_information/new", User_information.Post)
	router.HandleFunc("/user_information", User_information.Get)
	router.HandleFunc("/users/{user_uuid}", Users.GetOne)
	log.Fatal(http.ListenAndServe(":8008", router))

}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
