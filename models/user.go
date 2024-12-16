package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Roles    []string           `bson:"roles" json:"roles"`
}

type Profile struct {
	Id    primitive.ObjectID `bson:"_id" json:"_id"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email"`
	Roles []string           `bson:"roles" json:"roles"`
}

type InsertUser struct {
	Name     string   `bson:"name" json:"name"`
	Email    string   `bson:"email" json:"email"`
	Password string   `bson:"password" json:"password"`
	Roles    []string `bson:"roles" json:"roles"`
}

type UpdateUser struct {
	Id    string   `bson:"_id" json:"_id"`
	Name  string   `bson:"name" json:"name"`
	Email string   `bson:"email" json:"email"`
	Roles []string `bson:"roles" json:"roles"`
}
