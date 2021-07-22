package controllers

import (
	"net/http"
	"pebruwantoro/structuring/config"
	"pebruwantoro/structuring/lib/database"
	"pebruwantoro/structuring/models"
	"strconv"

	"github.com/labstack/echo"
)

// get all users
func GetUsersControllers(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

// get user by id
func GetUserControllers(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var count int64
	config.DB.Model(models.User{}).Where("id=?", id).Count(&count)
	if count == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	}
	getUser, err := database.GetOneUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get user by id",
		"users":  getUser,
	})
}

// create new user
func CreateUserControllers(c echo.Context) error {
	var user models.User
	c.Bind(&user)
	newUser, err := database.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success create new user",
		"users":  newUser,
	})
}

func DeleteUserControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	delete_user, err := database.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success delete user",
		"users":  delete_user,
	})
}

func UpdateUserControllers(c echo.Context) error {
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	get_user, _ := database.GetOneUser(id)
	user = get_user
	c.Bind(&user)
	update_user, err := database.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success update user",
		"users":  update_user,
	})
}
