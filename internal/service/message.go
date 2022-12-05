package service

import (
	"hte-device-update-consumer/internal/domain"
	"hte-device-update-consumer/internal/domain/gen"
	"hte-device-update-consumer/internal/repository"
)

type MessageService interface {
	RegisterLocation(m *gen.Message) error
}

type messageService struct {
	locationRepo repository.LocationRepository
	statusRepo   repository.StatusRepository
}

func NewMessageService(locationRepo repository.LocationRepository, statusRepo repository.StatusRepository) MessageService {
	return &messageService{
		locationRepo: locationRepo,
		statusRepo:   statusRepo,
	}
}

func (s *messageService) RegisterLocation(m *gen.Message) error {
	mDTO := domain.MessagePBToDTO(m)

	err := s.locationRepo.Create(mDTO)
	if err != nil {
		return err
	}

	err = s.statusRepo.Update(mDTO)

	return nil
}
