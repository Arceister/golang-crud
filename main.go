package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getuser", returnAllUsers).Methods("GET")
	router.HandleFunc("/registeruser", insertUsers).Methods("POST")
	router.HandleFunc("/updateuser", updateUsers).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Connected to port 69")
	log.Fatal(http.ListenAndServe(":69", router))
}
