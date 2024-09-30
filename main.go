package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"net/http"
)

// @title HelloWorld API
// @version 1.0
// @description No description
// @BasePath /
// @schemes http

// Message is the struct for the response containing a message
type Message struct {
	Message string `json:"message" validate:"required"`
}

// ErrorResponse is the struct for handling error responses
type ErrorResponse struct {
	Error string `json:"error" validate:"required"`
}

// PutModel is the struct for the PUT request
type PutModel struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/helloworld", getHelloWorld)   // GET
	e.PUT("/helloworld", putHelloWorld)   // PUT
	e.POST("/helloworld", postHelloWorld) // POST

	// Serve swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// getHelloWorld returns "hello world"
// @Summary getHelloWorld returns hello world
// @Description getHelloWorld returns hello world
// @Tags helloworld
// @Accept json
// @Produce json
// @Success 200 {object} Message
// @Failure 500 {object} ErrorResponse
// @Router /helloworld [get]
func getHelloWorld(c echo.Context) error {
	response := Message{
		Message: "Hello, World!",
	}
	return c.JSON(http.StatusOK, response)
}

// putHelloWorld returns "hello world" with the provided name
// @Summary getHelloWorld returns hello world with name
// @Description getHelloWorld returns hello world with name
// @Tags helloworld
// @Accept json
// @Produce json
// @Param name query string true "Name" minlength(2) maxlength(100) example(John)
// @Success 200 {object} Message
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /helloworld [put]
func putHelloWorld(c echo.Context) error {
	name := c.QueryParam("name")

	// Validate the input name
	if len(name) < 2 || len(name) > 100 {
		errResponse := ErrorResponse{
			Error: "Name must be between 2 and 100 characters long.",
		}
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	response := Message{
		Message: "Hello, " + name + "!",
	}
	return c.JSON(http.StatusOK, response)
}

// postHelloWorld accepts a name in the request body and returns a message
// @Summary getHelloWorld returns hello world with name
// @Description getHelloWorld returns hello world with name in the request body
// @Tags helloworld
// @Accept json
// @Produce json
// @Param query body PutModel true "what name"
// @Success 200 {object} Message
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /helloworld [post]
func postHelloWorld(c echo.Context) error {
	// Bind the request body to the PutModel struct
	putModel := new(PutModel)
	if err := c.Bind(putModel); err != nil {
		errResponse := ErrorResponse{
			Error: "Invalid input format",
		}
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	// Validate input
	if err := c.Validate(putModel); err != nil {
		errResponse := ErrorResponse{
			Error: "Invalid input: " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, errResponse)
	}

	response := Message{
		Message: "Hello, " + putModel.Name + "!",
	}
	return c.JSON(http.StatusOK, response)
}
