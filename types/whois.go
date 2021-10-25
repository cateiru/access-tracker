package types

import "time"

type People struct {
	Ip   string    `json:"ip"`
	Date time.Time `json:"time"`
}

type Whois struct {
	TrackId string   `json:"track_id"`
	History []People `json:"history"`
}
