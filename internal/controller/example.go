package controller

import (
	"github.com/gin-gonic/gin"
	"go-rest-template/internal/service"
)

type ExampleController interface {
	ExampleHandler(ctx *gin.Context)
}
type exampleController struct {
	svc service.ExampleService
}

func NewExampleController(svc service.ExampleService) ExampleController {
	return &exampleController{svc: svc}
}

func (c *exampleController) ExampleHandler(ctx *gin.Context) {

}
