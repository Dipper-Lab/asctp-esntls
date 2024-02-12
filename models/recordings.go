package models

import "time"

type Moisture struct {
	Value    float32   `bson:"value" json:"value"`
	Time     time.Time `bson:"time"  json:"time"`
	Location string    `bson:"location"  json:"location"`
}

type Note struct {
	Content string    `bson:"content" json:"content"`
	Bags    []string  `bson:"bags" json:"bags"`
	Time    time.Time `bson:"time" json:"time"`
}
