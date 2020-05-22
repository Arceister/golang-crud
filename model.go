package main

type Users struct {
	Id    string `form:"id" json:"id"`
	UName string `form:"uname" json:"username"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}

type Notify struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
