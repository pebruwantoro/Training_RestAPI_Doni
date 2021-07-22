package controllers

import (
	"net/http"
	"pebruwantoro/middleware/config"
	"pebruwantoro/middleware/lib/database"
	"pebruwantoro/middleware/middlewares"
	"pebruwantoro/middleware/models"
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

//jwt middlewares
func LoginUsersController(c echo.Context) error {
	userData := models.User{}
	c.Bind(&userData)
	users, err := database.LoginUsers(userData.Email, userData.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

func GetUserDetailControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// JWT Token from header
	loggedInUserId := middlewares.ExtractTokenUserId(c)
	if loggedInUserId != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized access, can not access the other databases")
	}
	users, err := database.GetDetailUsers((id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func DeleteOneUserControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	loggedInUserId := middlewares.ExtractTokenUserId(c)
	if loggedInUserId != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized access, can not access the other databases")
	}
	delete_user, err := database.DeleteOneUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  delete_user,
	})
}

func UpdateOneUserControllers(c echo.Context) error {
	var user models.User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	loggedInUserId := middlewares.ExtractTokenUserId(c)
	if loggedInUserId != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized access, can not access the other databases")
	}
	get_user, _ := database.GetDetailUsers(id)
	user = get_user
	c.Bind(&user)
	update_user, err := database.UpdateOneUser(user)
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
