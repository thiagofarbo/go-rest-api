package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-api/models"
	"user-api/repository"
	"user-api/services"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")

}

func Create(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	repository.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func Listar(w http.ResponseWriter, r *http.Request) {

	var p []models.User
	repository.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func FindById(w http.ResponseWriter, r *http.Request) {

	id := services.GetUserReference(r)

	var p models.User
	repository.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)

}

func Update(w http.ResponseWriter, r *http.Request) {

	id := services.GetUserReference(r)

	var user models.User

	repository.DB.First(&user, id)
	json.NewDecoder(r.Body).Decode(&user)

	repository.DB.Save(&user)
	json.NewEncoder(w).Encode(user)

}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := services.GetUserReference(r)

	var user models.User

	repository.DB.Delete(&user, id)
	json.NewEncoder(w).Encode(user)
}
