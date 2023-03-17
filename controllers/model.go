package controllers

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	UserType int    `json:"user_type"`
}

type UserResponse struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
	Data    User   `json:"Data"`
}

type UsersResponse struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
	Data    []User `json:"Data"`
}

type ErrorResponse struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
}
