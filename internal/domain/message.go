package domain

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
