package controller

import (
	"encoding/json"
	"hte-dispatcher/internal/domain"
	"hte-dispatcher/internal/service"
	"log"
)

type MessageController interface {
	Process(msg *string)
}
type messageController struct {
	svc service.MessageService
}

func NewMessageController(svc service.MessageService) MessageController {
	return &messageController{svc: svc}
}

func (c *messageController) Process(msg *string) {
	var m domain.MessageDTO
	bytes := []byte(*msg)
	err := json.Unmarshal(bytes, &m)

	if err != nil {
		log.Printf("Error unmarshaling payload: %+v\n", err)
		return
	}

	err = c.svc.RegisterLocation(&m)

	if err != nil {
		log.Printf("Error processing message: %+v\n", err)
	}
}
