package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:""`
	Name        string             `json:"room,omitempty"`
	PigstyNum   int                `json:"pigstynum,omitempty"`
	Temperature string             `json:"temperature,omitempty"`
	Humidity    string             `json:"humidity,omitempty"`
}

type Pig struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:""`
	Breed   string             `json:"breed,omitempty"`
	Dob     string             `json:"dob,omitempty"`
	Vaccine string             `json:"vaccine,omitempty"`
	Health  string             `json:"health,omitempty"`
}

type Pigsty struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:""`
	Type         string             `json:"pigstytype,omitempty"`
	Water_Bin_Id int                `json:"waterbin,omitempty"`
	Feed_Bin_Id  int                `json:"feedbin,omitempty"`
	Room_Id      int                `json:"room,omitempty"`
}
