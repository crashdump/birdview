package model

import "time"

type HealthEvent struct {
	Status      string `json:"status"`
	Name        string `json:"name"`
	Account     string `json:"account"`
	Region      string `json:"region"`
	Service     string `json:"service"`
	StartTime   time.Time `json:"start_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

type HealthEvents []HealthEvent
