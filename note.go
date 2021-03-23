package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Note adasfasdfasdf asdf asdfa s
type Note struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string
	Time time.Time
	Body string
}
