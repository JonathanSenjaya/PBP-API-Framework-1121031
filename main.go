package main

import (
	"fmt"

	"explorasiAPI/controllers"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()

	m.Group("/users", func(r martini.Router) {
		m.Delete("/:user_Id", controllers.DeleteUser)
		m.Get("", controllers.GetAllUserGorm)
		m.Post("", controllers.InsertUser)
		m.Put("/:user_Id", controllers.UpdatetUser)
	})

	//Latihan Week 7

	m.RunOnAddr(":8080")
	fmt.Printf("Connected to port 8080")
}
