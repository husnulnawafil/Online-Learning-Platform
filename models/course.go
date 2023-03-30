package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UUID      string             `bson:"uuid"`
	Name      string             `bson:"name"`
	Category  string             `bson:"category"`
	Price     float64            `bson:"price"`
	ImageURL  string             `bson:"image_url,omitempty"`
	IsFree    bool               `bson:"is_free"`
	Rating    float64            `bson:"rating,omitempty"`
	IsDeleted bool               `bson:"is_deleted,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UdpatedAt time.Time          `bson:"updated_at"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty"`
}
