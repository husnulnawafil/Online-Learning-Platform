package repositories

import (
	"context"
	"strings"
	"time"

	"github.com/husnulnawafil/online-learning-platform/configs/database"
	"github.com/husnulnawafil/online-learning-platform/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CourseRepositoryInstance struct {
	MongoDB    *mongo.Client
	Collection *mongo.Collection
}

func NewCourseRepository() *CourseRepositoryInstance {
	return &CourseRepositoryInstance{
		MongoDB:    database.MongoDb(),
		Collection: database.MongoDb().Database("test").Collection("courses"),
	}
}

func (c *CourseRepositoryInstance) CreateCourse(ctx context.Context, data models.Course) error {
	res, err := c.Collection.InsertOne(ctx, data)
	if res == nil {
		return err
	}
	return err
}

func (c *CourseRepositoryInstance) GetCourseByUUID(ctx context.Context, uuid string) (*models.Course, error) {
	final := new(models.Course)
	filter := bson.D{
		{Key: "uuid", Value: bson.D{
			{Key: "$eq", Value: uuid},
		}},
		{Key: "is_deleted", Value: bson.D{
			{Key: "$exists", Value: false},
		}},
	}
	res := c.Collection.FindOne(ctx, filter)
	res.Decode(&final)
	return final, nil
}

func (c *CourseRepositoryInstance) GetCourses(ctx context.Context, sort, filter primitive.D, search string) ([]models.Course, error) {
	var courses []models.Course
	if strings.TrimSpace(search) != "" {
		index := &mongo.IndexModel{Keys: bson.D{{Key: "name", Value: "text"}}}
		_, err := c.Collection.Indexes().CreateOne(context.TODO(), *index)
		if err != nil {
			return courses, err
		}
		searchFilter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: search}}}}
		filter = append(filter, searchFilter...)
	}
	opts := options.Find().SetSort(sort)
	cursor, err := c.Collection.Find(ctx, filter, opts)
	if err != nil {
		return courses, err
	}

	cursor.All(context.TODO(), &courses)
	if err != nil {
		return courses, err
	}

	return courses, err
}

func (c *CourseRepositoryInstance) UpdateCourse(ctx context.Context, UUID string, newData *models.Course) (int64, error) {
	filter := bson.D{{Key: "uuid", Value: UUID}}
	result, err := c.Collection.ReplaceOne(context.TODO(), filter, newData)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, err
}

func (c *CourseRepositoryInstance) DeleteCourse(ctx context.Context, UUID string) (int64, error) {
	filter := bson.D{{Key: "uuid", Value: UUID}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "is_deleted",
					Value: true,
				},
				{
					Key:   "deleted_at",
					Value: time.Now(),
				},
			},
		},
	}
	res, err := c.Collection.UpdateOne(context.TODO(), filter, update)

	return res.ModifiedCount, err
}
