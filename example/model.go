package example

import "time"

type User struct {
	Id        int64
	Name      string
	UserLabel Label
	At        time.Time
}

type Label struct {
	Id   int64
	Name string
}