package repositories

import (
	"context"
	"time"

	"github.com/husnulnawafil/online-learning-platform/configs/database"
	"github.com/husnulnawafil/online-learning-platform/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryInstance struct {
	MongoDB    *mongo.Client
	Collection *mongo.Collection
}

func NewUserRepository() *UserRepositoryInstance {
	return &UserRepositoryInstance{
		MongoDB:    database.MongoDb(),
		Collection: database.MongoDb().Database("test").Collection("users"),
	}
}

func (c *UserRepositoryInstance) RegisterUser(ctx context.Context, data models.User) error {
	res, err := c.Collection.InsertOne(ctx, data)
	if res == nil {
		return err
	}
	return err
}

func (c *UserRepositoryInstance) GetUserDetailByUUID(ctx context.Context, UUID string) (*models.User, error) {
	final := new(models.User)
	filter := bson.M{
		"uuid": UUID,
	}
	res := c.Collection.FindOne(ctx, filter)
	res.Decode(&final)
	return final, nil
}

func (c *UserRepositoryInstance) GetUserDetailByEmail(ctx context.Context, email string) (*models.User, error) {
	final := new(models.User)
	filter := bson.M{
		"email": email,
	}
	res := c.Collection.FindOne(ctx, filter)
	res.Decode(&final)
	return final, nil
}

func (c *UserRepositoryInstance) DeleteUser(ctx context.Context, UUID string) (int64, error) {
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
	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, err
}
