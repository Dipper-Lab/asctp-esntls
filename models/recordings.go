package models

import "time"

type Moisture struct {
	Value    float32   `bson:"value" json:"value"`
	Time     time.Time `bson:"time"  json:"time"`
	Location string    `bson:"location"  json:"location"`
}
