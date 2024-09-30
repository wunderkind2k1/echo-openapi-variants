package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// optional code omitted

type Server struct{}

func (s Server) PutHelloWorld(ctx echo.Context, params PutHelloWorldParams) error {
	resp := PutHelloWorldParams{
		Name: "Hello " + params.Name,
	}

	return ctx.JSON(http.StatusOK, resp)
}

func NewServer() ServerInterface {
	return Server{}
}

// (GET /helloworld)
func (Server) GetHelloWorld(ctx echo.Context) error {
	resp := PutHelloWorldParams{
		Name: "Hello World",
	}

	return ctx.JSON(http.StatusOK, resp)
}
