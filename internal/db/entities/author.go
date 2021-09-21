package entities

import "time"

type Author struct {
	Id    int64 `json:"ref"`
	Name  string
	Dated time.Time
}
