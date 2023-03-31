package services

import (
	"context"

	"github.com/husnulnawafil/online-learning-platform/models"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, data *models.CourseCategory) error
	GetCategoryDetail(ctx context.Context, name string) (*models.CourseCategory, error)
	GetCategories(ctx context.Context, sortBy, sortDir, search string) ([]models.CourseCategory, error)
}
