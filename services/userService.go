package services

import (
	"net/http"

	"fmt"
	"strconv"

	"github.com/gorilla/mux"
)

func SetContentType(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)

	})
}

func GetUserReference(r *http.Request) int {

	vars := mux.Vars(r)
	id := vars["id"]

	n, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Error during conversion")
		return 0
	}

	return n
}
