package services

import (
	"context"
	"strings"
	"time"

	"github.com/husnulnawafil/online-learning-platform/models"
	courseRepositories "github.com/husnulnawafil/online-learning-platform/repositories/courses"
	"go.mongodb.org/mongo-driver/bson"
)

type CourseServiceInstance struct {
	courseRepo courseRepositories.CourseRepository
}

func NewCourseService() CourseService {
	repoCourse := courseRepositories.NewCourseRepository()
	return &CourseServiceInstance{
		courseRepo: repoCourse,
	}
}

func (cs *CourseServiceInstance) CreateCourse(ctx context.Context, data *models.Course) error {
	data.CreatedAt = time.Now()
	data.UdpatedAt = time.Now()
	return cs.courseRepo.CreateCourse(ctx, *data)
}

func (cs *CourseServiceInstance) GetCourseByUUID(ctx context.Context, UUID string) (*models.Course, error) {
	return cs.courseRepo.GetCourseByUUID(ctx, UUID)
}

func (cs *CourseServiceInstance) GetCourses(ctx context.Context, sortBy, sortDir, search string, isFree bool) ([]models.Course, error) {
	filter := bson.D{
		{
			Key: "is_deleted", Value: bson.D{
				{Key: "$exists", Value: false},
			},
		},
	}

	filterIsFree := bson.D{
		{
			Key: "is_free", Value: bson.D{
				{Key: "$eq", Value: isFree},
			},
		},
	}

	if isFree {
		filter = append(filter, filterIsFree...)
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

	return cs.courseRepo.GetCourses(ctx, sort, filter, search)
}

func (cs *CourseServiceInstance) DeleteCourse(ctx context.Context, UUID string) (int64, error) {
	return cs.courseRepo.DeleteCourse(ctx, UUID)
}

func (cs *CourseServiceInstance) UpdateCourse(ctx context.Context, UUID string, newData *models.Course) (int64, error) {
	newData.UdpatedAt = time.Now()
	return cs.courseRepo.UpdateCourse(ctx, UUID, newData)
}
