package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func main() {
	e := echo.New()
	// routing with query
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	//start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

//---------------Controller----------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

//get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if users == nil {
		return c.JSON(http.StatusNotFound, "User Not Found")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by id",
		"users":    users[id-1],
	})
}

//delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	if users == nil {
		return c.JSON(http.StatusNotFound, "User Not Found")
	}
	users = append(users[:(id-1)], users[id:]...)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
		"users":    users[id],
	})
}

//update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newUser := User{}
	newUser = users[id]
	users = append(users[:(id-1)], users[id:]...)
	users = append(users, newUser)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user",
		"users":    newUser,
	})
}

//create new user
func CreateUserController(c echo.Context) error {
	//binding data
	user := User{}
	c.Bind(&user)
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}
