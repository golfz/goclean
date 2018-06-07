package output

import "time"

type Example struct {
	Key          string
	Name         string
	Color        string
	Remain       uint16
	CreatedDate  time.Time
}
