package repository

type ExampleRepository interface {
	ExampleFunction()
}

type exampleRepository struct {
}

func NewExampleRepository() ExampleRepository {
	return &exampleRepository{}
}

func (r *exampleRepository) ExampleFunction() {

}
