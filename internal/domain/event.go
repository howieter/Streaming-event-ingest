package domain

import "time"

type Event struct {
	UID     int            `json:"uid"`
	ID      string         `json:"id"`
	TS      time.Time      `json:"ts"`
	Type    string         `json:"type"`
	Source  string         `json:"source"`
	Payload map[string]any `json:"payload"`
}
