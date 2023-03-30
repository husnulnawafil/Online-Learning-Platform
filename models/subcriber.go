package models

type Subcriber struct {
	ID       uint   `json:"id"`
	UUID     string `json:"uuid"`
	UserID   uint   `json:"user_id"`
	CourseID uint   `json:"course_id"`
}
