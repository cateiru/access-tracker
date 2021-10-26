package types

import "time"

type History struct {
	Ip   string    `json:"ip"`
	Date time.Time `json:"time"`
}
