package Database

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string             `json:"email" bson:"email`
	Name  string             `json:"name" bson:"name,omitempty"`
	Pass  string             `json:"pass" bson:"pass,omitempty"`
	Notes []*Note            `json:"notes" bson:"notes,omitempty"`
}

type Note struct {
	Title     string `json:"title" bson:"title,omitempty"`
	Body      string `json:"body" bson:"body,omitempty"`
	Completed bool   `json:"completed" bson:"completed"`
}
