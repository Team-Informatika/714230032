package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Number    string             `bson:"number"`
	Message   string             `bson:"message"`
	Direction string             `bson:"direction"` // "incoming" atau "outgoing"
	Timestamp int64              `bson:"timestamp"`
}