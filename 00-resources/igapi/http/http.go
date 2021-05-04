package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Username      string `json:"username"`
	FullName      string `json:"full_name"`
	ProfilePicUrl string `json:"profile_pic_url"`
}

var allUsers = []User{
	{
		Username:      "golang",
		FullName:      "Gopher",
		ProfilePicUrl: "https://knowhere.fake/gopher.jpg",
	},
	{
		Username:      "golang-way",
		FullName:      "go away",
		ProfilePicUrl: "https://knowhere.fake/go-away.jpg",
	},
	{
		Username:      "python",
		FullName:      "Pythonistas",
		ProfilePicUrl: "https://knowhere.fake/python.jpg",
	},
	{
		Username:      "Javascript",
		FullName:      "Use me for everythings",
		ProfilePicUrl: "https://knowhere.fake/js.jpg",
	},
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("users handler ", r.Method, r.URL)
	if r.Method == http.MethodGet {
		var q = r.URL.Query()
		f := q.Get("filter")

		if f == "" {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(allUsers)
			return
		}

		var users []User
		for _, u := range allUsers {
			if strings.Contains(u.Username, f) {
				users = append(users, u)
			}
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)

	} else {
		msg := http.StatusText(http.StatusMethodNotAllowed)
		http.Error(w, msg, http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)

	port := "2021"
	log.Println("starting... port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
