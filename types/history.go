package types

import "time"

type History struct {
	UniqueId string    `json:"unique_id"`
	Ip       string    `json:"ip"`
	Date     time.Time `json:"time"`
}
