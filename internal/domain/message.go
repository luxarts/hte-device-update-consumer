package domain

import "hte-device-update-consumer/internal/domain/gen"

type MessageDTO struct {
	DeviceID  string      `json:"device_id"`
	Timestamp int64       `json:"ts"`
	Coords    Coordinates `json:"coords"`
	Battery   int64       `json:"battery"`
}

type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func MessagePBToDTO(m *gen.Message) *MessageDTO {
	return &MessageDTO{
		DeviceID:  m.DeviceID,
		Timestamp: m.Timestamp,
		Battery:   m.Battery,
		Coords: Coordinates{
			Latitude:  float64(m.Coordinates.Latitude),
			Longitude: float64(m.Coordinates.Longitude),
		},
	}
}
