package services

import (
	"context"

	"github.com/husnulnawafil/online-learning-platform/models"
)

type CourseService interface {
	GetCourseByUUID(ctx context.Context, UUID string) (*models.Course, error)
	CreateCourse(ctx context.Context, data *models.Course) error
	GetCourses(ctx context.Context, sortBy, orderType, search string, isFree bool) ([]models.Course, error)
	DeleteCourse(ctx context.Context, UUID string) (int64, error)
	UpdateCourse(ctx context.Context, UUID string, newData *models.Course) (int64, error)
}
