package handlers

import (
	"sync"

	services "github.com/husnulnawafil/online-learning-platform/services/categories"
)

var (
	service services.CategoryService
	once    sync.Once
)

func init() {
	once.Do(func() {
		service = services.NewCategoryService()
	})
}
