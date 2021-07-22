package routes

import (
	"pebruwantoro/structuring/controllers"

	"github.com/labstack/echo"
)

func New(e *echo.Echo) {

	//user
	e.GET("/users", controllers.GetUsersControllers)
	e.GET("/users/:id", controllers.GetUserControllers)
	e.POST("/users", controllers.CreateUserControllers)
	e.DELETE("/users/:id", controllers.DeleteUserControllers)
	e.PUT("/users/:id", controllers.UpdateUserControllers)

	//book
	e.GET("/books", controllers.GetBooksControllers)
	e.GET("/books/:id", controllers.GetBookControllers)
	e.POST("/books", controllers.CreateBookControllers)
	e.DELETE("/books/:id", controllers.DeleteBookControllers)
	e.PUT("/books/:id", controllers.UpdateBookControllers)

}
