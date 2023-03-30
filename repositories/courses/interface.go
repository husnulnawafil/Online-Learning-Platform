package repositories

import (
	"context"

	"github.com/husnulnawafil/online-learning-platform/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourseRepository interface {
	CreateCourse(ctx context.Context, data models.Course) error
	GetCourseByUUID(ctx context.Context, UUID string) (*models.Course, error)
	GetCourses(ctx context.Context, sortBy, orderType primitive.D, search string) ([]models.Course, error)
	DeleteCourse(ctx context.Context, UUID string) (int64, error)
	UpdateCourse(ctx context.Context, UUID string, newData *models.Course) (int64, error)
}
