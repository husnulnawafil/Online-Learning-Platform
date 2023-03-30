package repositories

import (
	"context"

	"github.com/husnulnawafil/online-learning-platform/configs/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StatisticRepositoryInstance struct {
	MongoDB    *mongo.Client
	CollCourse *mongo.Collection
	CollUser   *mongo.Collection
}

func NewStatisticRepository() *StatisticRepositoryInstance {
	return &StatisticRepositoryInstance{
		MongoDB:    database.MongoDb(),
		CollCourse: database.MongoDb().Database("test").Collection("courses"),
		CollUser:   database.MongoDb().Database("test").Collection("users"),
	}
}

func (s *StatisticRepositoryInstance) CountUser(ctx context.Context) (int64, error) {
	filter := bson.D{
		{
			Key: "is_deleted", Value: bson.D{
				{Key: "$exists", Value: false},
			},
		},
	}

	count, err := s.CollCourse.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, err
}

func (s *StatisticRepositoryInstance) CountCourse(ctx context.Context, isFree bool) (int64, error) {
	filter := bson.D{
		{
			Key: "is_deleted", Value: bson.D{
				{Key: "$exists", Value: false},
			},
		},
	}
	if isFree {
		freeFilter := bson.D{
			{
				Key: "is_free", Value: bson.D{
					{Key: "$eq", Value: isFree},
				},
			},
		}
		filter = append(filter, freeFilter...)
	}

	count, err := s.CollCourse.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, err
}
