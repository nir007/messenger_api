package application

import (
	"time"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Uid           string             `json:"uid" binding:"required"`
	Name          string             `json:"name" binding:"required"`
	SecondName    string             `json:"second" binding:"required"`
	Avatar        string             `json:"avatar"`
	Gender        string             `json:"gender" binding:"required"`
	Links         []string           `json:"links"`
	Email         string             `json:"email" binding:"omitempty,email"`
	Phone         string             `json:"phone"`
	BlackList     []string           `json:"blackList"`
	CreatedAt     string             `json:"createdAt" binding:"-"`
	UpdatedAt     string             `json:"updatedAt" binding:"-"`
	ApplicationID string             `json:"applicationID" binding:"required"`
}

func (u *User) Insert() (*mongo.InsertOneResult, error) {
	if err := u.FindOne(bson.M{"uid": u.Uid}); err == nil {
		return nil, fmt.Errorf("user with id: %s already exists", u.Uid)
	}

	collection := client.Database("messenger").Collection("users_" + u.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now().String()
	u.UpdatedAt = ""

	return collection.InsertOne(ctx, u)
}

func (u *User) Find(c *gin.Context, find MongoParamsGetter) ([]User, int64, error) {
	result := make([]User, 0)

	collection := client.Database("messenger").Collection("users_" + u.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	findBson := find.ToBson()
	params := &options.FindOptions{Skip: find.Skip(), Limit: find.Limit()}

	total, err := collection.CountDocuments(ctx, findBson)
	if err != nil {
		return  result, 0, err
	}

	cur, err := collection.Find(ctx, findBson, params)
	if err != nil {
		return  result, 0, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		u := &User{}
		err := cur.Decode(u)
		result = append(result, *u)
		if err != nil {
			return  result, 0, err
		}
	}

	return result, total, err
}

func (u *User) FindOne(find bson.M) error {
	collection := client.Database("messenger").Collection("users_" + u.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	return collection.FindOne(ctx, find).Decode(u)
}
