package controller

import (
	"hte-device-update-consumer/internal/domain/gen"
	"hte-device-update-consumer/internal/service"
	"log"

	"google.golang.org/protobuf/proto"
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
	var m gen.Message

	err := proto.Unmarshal([]byte(*msg), &m)
	if err != nil {
		log.Printf("Error unmarshaling payload: %+v\n", err)
		return
	}
	log.Println(m)

	err = c.svc.RegisterLocation(&m)

	if err != nil {
		log.Printf("Error processing message: %+v\n", err)
	}
}
