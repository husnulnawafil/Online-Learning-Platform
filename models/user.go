package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UUID         string             `bson:"uuid"`
	Name         string             `bson:"name"`
	ProfileImage string             `bson:"profile_image,omitempty"`
	Email        string             `bson:"email"`
	Password     string             `bson:"password"`
	Role         string             `bson:"role"`
	IsDeleted    bool               `bson:"is_deleted,omitempty"`
	CreatedAt    time.Time          `bson:"created_at"`
	UdpatedAt    time.Time          `bson:"updated_at"`
	DeletedAt    time.Time          `bson:"deleted_at,omitempty"`
}
