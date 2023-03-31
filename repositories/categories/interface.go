package repositories

import (
	"context"

	"github.com/husnulnawafil/online-learning-platform/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, data models.CourseCategory) error
	GetCategoryDetail(ctx context.Context, name string) (*models.CourseCategory, error)
	GetCategories(ctx context.Context, sort, filter primitive.D, search string) ([]models.CourseCategory, error)
}
