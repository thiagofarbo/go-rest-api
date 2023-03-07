package main

import (
	"fmt"
	"user-api/repository"
	"user-api/routes"
)

func main() {

	// models.Users = []models.User{
	// 	{
	// 		Id: 1, Name: "Thiago", Address: "Rua 1",
	// 	},
	// 	{
	// 		Id: 2, Name: "Emidio", Address: "Rua 2",
	// 	},
	// 	{
	// 		Id: 3, Name: "Correa", Address: "Rua 3",
	// 	},
	// }
	repository.ConnectDatabase()
	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
