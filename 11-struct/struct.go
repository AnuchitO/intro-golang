package main

import "fmt"

type User struct {
	Username      string
	FullName      string
	ProfilePicUrl string
}

func main() {
	username := "golang"
	fullName := "Basic Golang"
	profilePicUrl := "https://knowhere.fake/gopher.jpg"
	fmt.Println(username, fullName, profilePicUrl)

	u := User{
		Username:      username,
		FullName:      fullName,
		ProfilePicUrl: profilePicUrl,
	}

	// u := User{}
	// u.Username = username
	// u.FullName = fullName
	// u.ProfilePicUrl = profilePicUrl

	fmt.Printf("%#v\n", u)

	fmt.Println(u.Username)
	fmt.Println(u.FullName)
	fmt.Println(u.ProfilePicUrl)

}
