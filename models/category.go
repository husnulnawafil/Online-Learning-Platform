package models

import "time"

type CourseCategory struct {
	Name           string    `bson:"name" json:"name"`
	TotalSubcriber int64     `bson:"total_subcriber,omitempty"`
	AverageRating  float32   `bson:"average_rating,omitempty"`
	IsDeleted      bool      `bson:"is_deleted,omitempty"`
	CreatedAt      time.Time `bson:"created_at"`
	UdpatedAt      time.Time `bson:"updated_at"`
	DeletedAt      time.Time `bson:"deleted_at,omitempty"`
}
