package services

import (
	"context"
	"strings"
	"time"

	"github.com/husnulnawafil/online-learning-platform/models"
	categoryRepositories "github.com/husnulnawafil/online-learning-platform/repositories/categories"
	"go.mongodb.org/mongo-driver/bson"
)

type CategoryServiceInstance struct {
	categoryRepo categoryRepositories.CategoryRepository
}

func NewCategoryService() CategoryService {
	repoCategory := categoryRepositories.NewCategoryRepository()
	return &CategoryServiceInstance{
		categoryRepo: repoCategory,
	}
}

func (cs *CategoryServiceInstance) CreateCategory(ctx context.Context, data *models.CourseCategory) error {
	data.CreatedAt = time.Now()
	data.UdpatedAt = time.Now()
	return cs.categoryRepo.CreateCategory(ctx, *data)
}

func (cs *CategoryServiceInstance) GetCategoryDetail(ctx context.Context, name string) (*models.CourseCategory, error) {
	return cs.categoryRepo.GetCategoryDetail(ctx, name)
}

func (cs *CategoryServiceInstance) GetCategories(ctx context.Context, sortBy, sortDir, search string) ([]models.CourseCategory, error) {
	filter := bson.D{
		{
			Key: "is_deleted", Value: bson.D{
				{Key: "$exists", Value: false},
			},
		},
	}

	sortDirection := 1
	if sortDir == strings.ToLower("desc") {
		sortDirection = -1
	}

	sort := bson.D{
		{
			Key:   sortBy,
			Value: sortDirection,
		},
	}

	return cs.categoryRepo.GetCategories(ctx, sort, filter, search)
}
