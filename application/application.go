package application

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	_"go.mongodb.org/mongo-driver/bson/primitive"
)

type managers struct {
	ID          string
	Name        string `validate:"required"`
	Surname     string `validate:"required"`
	Email       string `validate:"required"`
	Password    string `validate:"required,email"`
	LastLoginAt string
	CreatedAt   string
	UpdatedAt   string
	Group       []string
	Role        string
}

type Application struct {
	ID          string     `json:"id"`
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description" binding:"required"`
	Secret      string     `json:"secret"`
	Domains     []string   `json:"domains" binding:"required"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	Managers    []managers `json:"managers" binding:"required"`
}

func (a *Application) Insert() (*mongo.InsertOneResult, error) {
	collection := client.Database("messenger").Collection("applications")
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	u1 := uuid.Must(uuid.NewV4(), nil)
	a.Secret = fmt.Sprintf("%s", u1)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, a)

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		a.ID = oid.String()
	}

	return res, err
}

func (a *Application)FindOne() error {
	collection := client.Database("messenger").Collection("applications")

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"ID": a.ID}

	return collection.FindOne(ctx, filter).Decode(a)
}
