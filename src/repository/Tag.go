package repository


import "time"

type Tag struct {
	Base
	Name 			string
	Architecture 	string
	LastUsed		time.Time
}