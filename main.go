package main

import (
	"fmt"
	"net/http"
	"rest_api/users"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from go server")
}
func main() {
	http.HandleFunc("/update", users.UpdateUser)
	http.HandleFunc("/delete", users.DeleteUser)
	http.HandleFunc("/user", users.GetUser)
	http.HandleFunc("/users", users.CreateUser)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/ping", users.Ping)
	fmt.Println("Server starting on :8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
