package routes

import (
	"pebruwantoro/middleware/constants"
	"pebruwantoro/middleware/controllers"
	"pebruwantoro/middleware/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	//user
	e.GET("/users", controllers.GetUsersControllers)
	e.GET("/users/:id", controllers.GetUserControllers)
	e.POST("/users", controllers.CreateUserControllers)
	e.DELETE("/users/:id", controllers.DeleteUserControllers)
	e.PUT("/users/:id", controllers.UpdateUserControllers)

	//login to get token
	e.POST("/login", controllers.LoginUsersController)

	// JWT_Middlewares
	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users/:id", controllers.GetUserDetailControllers)
	eJwt.DELETE("/users/:id", controllers.DeleteOneUserControllers)
	eJwt.PUT("/users/:id", controllers.UpdateOneUserControllers)

	//book
	e.GET("/books", controllers.GetBooksControllers)
	e.GET("/books/:id", controllers.GetBookControllers)
	e.POST("/books", controllers.CreateBookControllers)
	e.DELETE("/books/:id", controllers.DeleteBookControllers)
	e.PUT("/books/:id", controllers.UpdateBookControllers)

	// eauthGroup
	eAuth := e.Group("")
	eAuth.Use(middleware.BasicAuth(middlewares.BasicAuthDb))
}
