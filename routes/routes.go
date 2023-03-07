package routes

import (
	"log"
	"net/http"
	"user-api/controllers"
	"user-api/services"

	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()

	r.Use(services.SetContentType)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/users", controllers.Listar).Methods("Get")
	r.HandleFunc("/api/users/{id}", controllers.FindById).Methods("Get")
	r.HandleFunc("/api/users", controllers.Create).Methods("Post")
	r.HandleFunc("/api/users/{id}", controllers.Delete).Methods("Delete")
	r.HandleFunc("/api/users/{id}", controllers.Update).Methods("Put")
	log.Fatal(http.ListenAndServe(":8091", r))
}
