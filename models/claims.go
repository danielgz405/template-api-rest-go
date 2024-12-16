package models

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AppClaims struct {
	UserId primitive.ObjectID `bson:"userId"`
	jwt.StandardClaims
}
