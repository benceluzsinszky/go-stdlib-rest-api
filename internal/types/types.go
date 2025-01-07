package types

import (
	"time"
)

type Item struct {
	Id   int64     `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}
