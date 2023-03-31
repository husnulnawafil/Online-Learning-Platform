package repositories

import (
	"context"
	"strings"

	"github.com/husnulnawafil/online-learning-platform/configs/database"
	"github.com/husnulnawafil/online-learning-platform/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryRepositoryInstance struct {
	MongoDB    *mongo.Client
	Collection *mongo.Collection
}

func NewCategoryRepository() *CategoryRepositoryInstance {
	return &CategoryRepositoryInstance{
		MongoDB:    database.MongoDb(),
		Collection: database.MongoDb().Database("test").Collection("categories"),
	}
}

func (c *CategoryRepositoryInstance) CreateCategory(ctx context.Context, data models.CourseCategory) error {
	res, err := c.Collection.InsertOne(ctx, data)
	if res == nil {
		return err
	}
	return err
}

func (c *CategoryRepositoryInstance) GetCategoryDetail(ctx context.Context, name string) (*models.CourseCategory, error) {
	final := new(models.CourseCategory)
	filter := bson.D{
		{Key: "name", Value: bson.D{
			{Key: "$eq", Value: name},
		}},
		{Key: "is_deleted", Value: bson.D{
			{Key: "$exists", Value: false},
		}},
	}
	res := c.Collection.FindOne(ctx, filter)
	res.Decode(&final)
	return final, nil
}

func (c *CategoryRepositoryInstance) GetCategories(ctx context.Context, sort, filter primitive.D, search string) ([]models.CourseCategory, error) {
	var categories []models.CourseCategory
	if strings.TrimSpace(search) != "" {
		index := &mongo.IndexModel{Keys: bson.D{{Key: "name", Value: "text"}}}
		_, err := c.Collection.Indexes().CreateOne(context.TODO(), *index)
		if err != nil {
			return categories, err
		}
		searchFilter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: search}}}}
		filter = append(filter, searchFilter...)
	}
	opts := options.Find().SetSort(sort)
	cursor, err := c.Collection.Find(ctx, filter, opts)
	if err != nil {
		return categories, err
	}

	cursor.All(context.TODO(), &categories)
	if err != nil {
		return categories, err
	}

	return categories, err
}
