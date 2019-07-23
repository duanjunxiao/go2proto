package project

import (
	"time"
)

type Label struct {
	Id   int64
	Name string
}

type User struct {
	Id   int64
	Tag  Label
	List []Label
	At   time.Time
}
