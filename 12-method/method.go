package main

import "fmt"

type User struct {
	Username      string
	FullName      string
	ProfilePicUrl string
}

func (u User) Info() {
	fmt.Printf("User name : %v\n", u.Username)
	fmt.Printf("Full name : %v\n", u.FullName)
	fmt.Printf("Profile Picture URL : %v\n", u.ProfilePicUrl)
}

func main() {
	u := User{
		Username:      "golang",
		FullName:      "Basic Golang",
		ProfilePicUrl: "https://knowhere.fake/gopher.jpg",
	}

	// info(u)
	u.Info()
}
