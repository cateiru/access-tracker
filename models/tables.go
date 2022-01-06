package models

import "time"

type Track struct {
	TrackId     string    `datastore:"trackId"`
	AccessKey   string    `datastore:"accessKey"`
	RedirectUrl string    `datastore:"redirectUrl"`
	Create      time.Time `datastore:"create"`
}

type History struct {
	UniqueId  string    `datastore:"uniqueId" json:"unique_id"`
	TrackId   string    `datastore:"trackId" json:"track_id"`
	Ip        string    `datastore:"ip" json:"ip"`
	UserAgent string    `datastore:"datastore" json:"datastore"`
	Date      time.Time `datastore:"date" json:"time"`
}
