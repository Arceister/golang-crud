package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arr_user []Users
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, uname FROM users")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.UName); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_user = append(arr_user, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func insertUsers(w http.ResponseWriter, r *http.Request) {
	var response Notify

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	uname := r.FormValue("username")

	_, err = db.Exec("INSERT INTO users (uname) VALUES (?)", uname)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "User Registered!"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updateUsers(w http.ResponseWriter, r *http.Request) {
	var response Notify

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("uid")
	uname := r.FormValue("username")

	_, err = db.Exec("UPDATE users SET uname = ? WHERE id = ?", uname, id)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "User updated!"
	log.Print("DB updated!")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
