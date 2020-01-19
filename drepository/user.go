package drepository

import (
	"context"
	"fmt"
	"messenger/dto"
	"time"

	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User struct for messenger app users
type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UID           string             `json:"uid" binding:"required"`
	ApplicationID string             `json:"applicationID" binding:"required"`
	Name          string             `json:"name" binding:"required"`
	SecondName    string             `json:"second" binding:"required"`
	Avatar        string             `json:"avatar"`
	Gender        string             `json:"gender" binding:"required"`
	Links         []string           `json:"links"`
	Email         string             `json:"email" binding:"omitempty,email"`
	Phone         string             `json:"phone"`
	About         string             `json:"about" binding:"omitempty,email"`
	BlackList     []string           `json:"blackList"`
	CreatedAt     string             `json:"createdAt" binding:"-"`
	UpdatedAt     string             `json:"updatedAt" binding:"-"`
	DeletedAt     string             `json:"deletedAt" binding:"-"`
}

// Delete deletes documents
func (mc *User) Delete() (int64, error) {
	collection := client.Database(dbName).Collection("users_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	updated, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": mc.ID},
		bson.M{"$set": bson.M{"deletedat": time.Now().String()}},
	)

	if err != nil {
		return 0, err
	}

	return updated.ModifiedCount, err
}

// Update updates one document
func (mc *User) Update(find dto.SearchParamsGetter, update dto.BSONMaker) (int64, error) {
	collection := client.Database(dbName).Collection("users_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	updateResult, err := collection.UpdateOne(
		ctx,
		find.ToBson(),
		bson.M{"$set": update.ToBson()},
	)

	if err != nil {
		return 0, err
	}

	if updateResult == nil || updateResult.ModifiedCount == 0 {
		return 0, errors.New("undefined user")
	}

	_ = mc.FindOne(find)
	return updateResult.ModifiedCount, nil
}

//Insert creates new document
func (mc *User) Insert() (string, error) {
	find := &dto.FindUsers{UID: mc.UID}

	if err := mc.FindOne(find); err == nil {
		return "", fmt.Errorf("user with UID: %s already exists", find.UID)
	}

	collection := client.Database(dbName).Collection("users_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	mc.ID = primitive.NewObjectID()
	mc.CreatedAt = time.Now().String()
	mc.UpdatedAt = ""

	res, err := collection.InsertOne(ctx, mc)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", res.InsertedID), err
}

// FindOne finds one document
func (mc *User) FindOne(find dto.SearchParamsGetter) error {
	collection := client.Database(dbName).Collection("users_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	return collection.FindOne(ctx, find.ToBson()).Decode(mc)
}

// Find finds several documents by pages
func (mc *User) Find(find dto.SearchParamsGetter) ([]interface{}, int64, error) {
	result := make([]interface{}, 0)

	collection := client.Database(dbName).Collection("users_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	findBson := find.ToBson()
	options := &options.FindOptions{Skip: find.Skip(), Limit: find.Limit(), Sort: find.Sort()}

	total, err := collection.CountDocuments(ctx, findBson)
	if err != nil {
		return result, 0, err
	}

	cur, err := collection.Find(ctx, findBson, options)
	if err != nil {
		return result, 0, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		u := &User{}
		err := cur.Decode(u)
		result = append(result, *u)
		if err != nil {
			return result, 0, err
		}
	}

	return result, total, err
}
