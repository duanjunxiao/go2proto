package in

import "time"

type Entity struct {
	Kind  int64
	Value string
}

type Message struct {
	Entity    Entity
	CreatedAt time.Time
}
