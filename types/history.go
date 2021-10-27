package types

import "time"

type History struct {
	UniqueId string    `datastore:"UniqueId" json:"unique_id"`
	Ip       string    `datastore:"Ip" json:"ip"`
	Date     time.Time `datastore:"Date" json:"time"`
}
