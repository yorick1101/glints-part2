package model

import "time"

type IntFilter struct {
	Operation string //eq, lt, gt
	Value     int
}

type DateFilter struct {
	Operation string //eq, lt, gt
	Value     time.Time
}
