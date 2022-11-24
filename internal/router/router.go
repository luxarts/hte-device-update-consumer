package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luxarts/jsend-go"
	"go-rest-template/internal/controller"
	"go-rest-template/internal/defines"
	"go-rest-template/internal/repository"
	"go-rest-template/internal/service"
	"net/http"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// DB connectors, rest clients, and other stuff init

	// Repositories init
	repo := repository.NewExampleRepository()

	// Services init
	svc := service.NewExampleService(repo)

	// Controllers init
	ctrl := controller.NewExampleController(svc)

	// Endpoints
	r.GET(defines.EndpointExample, ctrl.ExampleHandler)

	// Health check endpoint
	r.GET(defines.EndpointPing, healthCheck)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, jsend.NewSuccess("pong"))
}
