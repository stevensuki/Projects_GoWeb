package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	e := echo.New()

	users := e.Group("/users")
	users.GET("/", sayHello)
	users.GET("/:id", sayHelloParam)
	users.GET("/query", sayHelloQueryParam)
	users.GET("/json", sayHellomJSON)

	e.Logger.Fatal(e.Start(":8080"))
}

func sayHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func sayHelloParam(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Hello "+id)
}

func sayHelloQueryParam(c echo.Context) error {
	firstName := c.QueryParam("firstname")
	lastName := c.QueryParam("lastname")
	return c.String(http.StatusOK, "Hello "+firstName+" "+lastName)
}

func sayHellomJSON(c echo.Context) error {
	return c.JSON(http.StatusOK, User{Firstname: "Steven", Lastname: "Suki"})
}
