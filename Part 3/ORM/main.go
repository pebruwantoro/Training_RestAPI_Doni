package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "02021996Doni*",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DB_Username,
			config.DB_Password,
			config.DB_Host,
			config.DB_Port,
			config.DB_Name,
		)
	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"email" form:"email"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

func GetUsersController(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success get all users",
		"users":   users,
	})
}

// getting user by id
func GetUserController(c echo.Context) error {
	var user []User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var count int64
	DB.Model(User{}).Where("id=?", id).Count(&count)
	if count == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	}
	if err := DB.Find(&user, "id=?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by id",
		"users":    user,
	})
}

// creating new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	if err := DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			":message": "can not insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// deleting user
func DeleteUserController(c echo.Context) error {
	var user []User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	if err := DB.Find(&user, "id=?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			":message": "can not fetch data",
		})
	}
	if err := DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			":message": "can not post data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "success delete user",
		"users":   user,
	})
}

// updating user data

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	var user User
	if err := DB.Find(&user, "id=?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			":message": "can not fetch data",
		})
	}
	c.Bind(&user)
	if err := DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			":message": "can not update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user",
		"users":    user,
	})
}

func main() {
	//create a new echo instance
	e := echo.New()

	//Route / to handler function
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	//start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
