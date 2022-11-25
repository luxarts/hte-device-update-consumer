package service

import (
	"hte-device-update-consumer/internal/domain"
	"hte-device-update-consumer/internal/repository"
)

type MessageService interface {
	RegisterLocation(msg *domain.MessageDTO) error
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

func (s *messageService) RegisterLocation(msg *domain.MessageDTO) error {
	err := s.locationRepo.Create(msg)
	if err != nil {
		return err
	}

	err = s.statusRepo.Update(msg)

	return nil
}
