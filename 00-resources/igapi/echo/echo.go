package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Username      string `json:"username"`
	FullName      string `json:"full_name"`
	ProfilePicUrl string `json:"profile_pic_url"`
}

var u1 = User{
	Username:      "golang",
	FullName:      "Gopher",
	ProfilePicUrl: "https://knowhere.fake/gopher.jpg",
}
var u2 = User{
	Username:      "golang-way",
	FullName:      "go away",
	ProfilePicUrl: "https://knowhere.fake/go-away.jpg",
}

var u3 = User{
	Username:      "python",
	FullName:      "Pythonistas",
	ProfilePicUrl: "https://knowhere.fake/python.jpg",
}
var u4 = User{
	Username:      "Javascript",
	FullName:      "Use me for everythings",
	ProfilePicUrl: "https://knowhere.fake/js.jpg",
}

var allUsers = []User{u1, u2, u3, u4}

func usersHandler(c echo.Context) error {
	f := c.QueryParam("filter")
	if f == "" {
		err := c.JSON(http.StatusOK, allUsers)
		return err
	}

	var users []User
	for _, u := range allUsers {
		if strings.Contains(u.Username, f) {
			users = append(users, u)
		}
	}

	return c.JSON(http.StatusOK, users)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/users", usersHandler)

	port := "2021"
	log.Println("starting... port:", port)
	log.Fatal(e.Start(":" + port))
}
