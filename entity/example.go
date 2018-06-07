package entity

import (
	"time"
)

type ExampleEntity struct {
	Key          string
	Name         string
	Color        string
	Remain       uint16
	HasPromotion bool
	CreatedDate  time.Time
	Size         int32
	Type         uint8
}
