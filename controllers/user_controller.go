package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
)

func GetAllUserGorm(w http.ResponseWriter, r *http.Request) {
	db := ConnectGorm()

	var users []User
	db.Find(&users)
	w.Header().Set("content-Type", "application/json")
	var response UsersResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = users
	json.NewEncoder(w).Encode(response)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	// Read from request body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")
	userType, _ := strconv.Atoi(r.Form.Get("age"))

	_, errQuery := db.Exec("INSERT INTO users(name, age, address, type) VALUES (?,?,?,?)",
		name,
		age,
		address,
		userType,
	)

	var user User
	user.Name = name
	user.Age = age
	user.Address = address
	user.UserType = userType
	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		response.Data = user
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "insert Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}
	userId := params["user_id"]

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?",
		userId,
	)

	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Delete Failed!"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdatetUser(params martini.Params, w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()
	userId := params["user_id"]

	// Read from request body
	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")

	_, errQuery := db.Exec("UPDATE users SET name=?, age=?, address=? WHERE id=?",
		name,
		age,
		address,
		userId,
	)

	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		response.Data.Name = name
		response.Data.Address = address
		response.Data.Age = age
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "insert Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	var response ErrorResponse
	response.Status = 400
	response.Message = message
	//response.Err = err
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
