package entity

import "time"

type Size struct {
	SizeRequest float64   `json:"sizeRequest"`
	Time        time.Time `json:"time"`
}
