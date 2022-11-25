package repository

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"hte-device-update-consumer/internal/defines"
	"hte-device-update-consumer/internal/domain"
	"log"
	"net/http"
	"os"
)

type StatusRepository interface {
	Update(status *domain.MessageDTO) error
}

type statusRepository struct {
	baseURL string
	client  *resty.Client
}

func NewStatusRepository(client *resty.Client) StatusRepository {
	return &statusRepository{
		baseURL: os.Getenv(defines.EnvAPIStatusHost),
		client:  client,
	}
}

func (r *statusRepository) Update(status *domain.MessageDTO) error {
	req := r.client.R()
	req = req.SetBody(status)
	req = req.SetPathParam(defines.ParamDeviceID, status.DeviceID)
	resp, err := req.Put(r.baseURL + defines.APIPutStatus)

	if err != nil {
		return err
	}

	log.Printf("Status update -> %d\n", resp.StatusCode())

	if resp.StatusCode() != http.StatusOK {
		return errors.New("location service failed")
	}

	return nil
}
