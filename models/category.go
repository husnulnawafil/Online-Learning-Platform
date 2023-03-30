package models

type CourseCategory struct {
	Name      string `bson:"name"`
	Subcriber int64  `bson:"subcriber"`
}
