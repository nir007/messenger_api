package drepository

import (
	"context"
	"fmt"
	"messenger/dto"
	"time"
	"errors"

	uuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Application struct for messeneger application
type Application struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" binding:"required,max=50,min=1"`
	Description string             `json:"description" binding:"required,max=100,min=1"`
	Secret      string             `json:"secret"`
	Domains     []string           `json:"domains" binding:"required"`
	CreatedAt   string             `json:"createdAt" bson:"-"`
	UpdatedAt   string             `json:"updatedAt" bson:"-"`
	DeletedAt   string             `json:"deletedAt" binding:"-"`
	Managers    []string           `json:"managers" binding:"required"`
}

// Delete deletes documents
func (mc *Application) Delete(id, applicationID string) (int64, error) {
	collection := client.Database(dbName).Collection("applications")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	deleteResult, err := collection.DeleteOne(ctx, bson.M{"_id": mc.ID})

	return deleteResult.DeletedCount, err
}

// Update updates documents
func (mc *Application) Update(find dto.SearchParamsGetter, update dto.BSONMaker) (int64, error) {
	collection := client.Database(dbName).Collection("applications")
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
		return 0, errors.New("undefined application")
	}

	_ = mc.FindOne(find)
	return updateResult.ModifiedCount, err
}

//Insert creates new document
func (mc *Application) Insert() (string, error) {
	collection := client.Database(dbName).Collection("applications")

	mc.ID = primitive.NewObjectID()
	mc.CreatedAt = time.Now().String()
	mc.UpdatedAt = ""
	mc.Secret = uuid.New().String()

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, mc)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", res.InsertedID), err
}

// FindOne finds one document
func (mc *Application) FindOne(find dto.SearchParamsGetter) error {
	collection := client.Database(dbName).Collection("applications")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	return collection.FindOne(ctx, find.ToBson()).Decode(mc)
}

// Find finds several documents by pages
func (mc *Application) Find(find dto.SearchParamsGetter) ([]interface{}, int64, error) {
	result := make([]interface{}, 0)

	collection := client.Database(dbName).Collection("applications")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	opt := &options.FindOptions{Skip: find.Skip(), Limit: find.Limit(), Sort: find.Sort()}

	total, err := collection.CountDocuments(ctx, find.ToBson())
	if err != nil {
		return result, 0, err
	}

	cur, err := collection.Find(ctx, find.ToBson(), opt)
	if err != nil {
		return make([]interface{}, 0), 0, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		app := &Application{}
		err := cur.Decode(app)
		result = append(result, *app)
		if err != nil {
			return make([]interface{}, 0), 0, errors.New("dick pussy")
		}
	}

	return result, total, nil
}
