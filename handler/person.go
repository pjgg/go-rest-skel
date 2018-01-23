package handler

import (
	"net/http"

	cErrors "github.com/pjgg/go-errors/errors"

	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	errorHandler cErrors.ErrorHandlerBehavior
}

type PersonHandlerBehavior interface {
	// swagger:route GET /v1/person/ping ping example
	//
	// ping method
	//
	// This will always return pong
	// second line of description
	//
	// Consumes:
	// application/json
	//
	// Produces:
	// application/json
	//
	// Schemes: http
	//
	// Responses:
	// 200: someResponse
	// swagger:meta
	Ping()
}

func NewPersonHandler(errorHandler cErrors.ErrorHandlerBehavior) *PersonHandler {
	return &PersonHandler{errorHandler: errorHandler}
}

func (p *PersonHandler) setupPersonAPIRouting(g *gin.RouterGroup) {
	g.GET("/ping", p.Ping())
}

func (p *PersonHandler) Ping() func(c *gin.Context) {
	return func(c *gin.Context) {
		//requestID := c.Keys[utils.Uid].(string)

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
