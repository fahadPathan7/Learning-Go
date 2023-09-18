package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// this is the model for the data we are going to store in mongodb
type Netflix struct {
	// primitive.ObjectID is the unique id for the record in mongodb. It is a 12-byte BSON type. It is required by mongodb.
	// omitempty means that if the field has empty value then it will be omitted from the encoded JSON.
	// bson is the key for the field in mongodb. It is required by mongodb.

	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie string `json:"movie,omitempty" bson:"movie,omitempty"`
	Watched bool `json:"watched,omitempty" bson:"watched,omitempty"`
}