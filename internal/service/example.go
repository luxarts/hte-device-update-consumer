package service

import "go-rest-template/internal/repository"

type ExampleService interface {
	ExampleFunction()
}

type exampleService struct {
	repo repository.ExampleRepository
}

func NewExampleService(repo repository.ExampleRepository) ExampleService {
	return &exampleService{repo: repo}
}

func (r *exampleService) ExampleFunction() {

}
