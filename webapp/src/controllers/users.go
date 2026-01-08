package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// UserRegister calls an API to register a user in the database
func UserRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, error := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if error != nil {
		log.Fatal(error)
	}

	response, error := http.Post("http://localhost:5000/users", "application/json", bytes.NewBuffer(user))
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()

	fmt.Println(response.Body)
}
