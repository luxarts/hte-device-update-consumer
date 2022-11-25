package repository

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"hte-dispatcher/internal/defines"
	"hte-dispatcher/internal/domain"
	"log"
	"net/http"
	"os"
)

type LocationRepository interface {
	Create(location *domain.MessageDTO) error
}

type locationRepository struct {
	baseURL string
	client  *resty.Client
}

func NewLocationRepository(client *resty.Client) LocationRepository {
	return &locationRepository{
		baseURL: os.Getenv(defines.EnvAPILocationHost),
		client:  client,
	}
}

func (r *locationRepository) Create(location *domain.MessageDTO) error {
	req := r.client.R()
	req = req.SetBody(location)
	resp, err := req.Post(r.baseURL + defines.APIPostLocation)

	if err != nil {
		return err
	}

	log.Printf("Location create -> %d\n", resp.StatusCode())

	if resp.StatusCode() != http.StatusCreated {
		return errors.New("location service failed")
	}

	return nil
}
