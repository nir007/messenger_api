package application

import (
	"context"
	"fmt"
	"messenger/dto"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Application struct for messeneger application
type Application struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" binding:"required,max=50,min=1"`
	Description string             `json:"description" binding:"required,max=100,min=1"`
	Secret      string             `json:"secret"`
	Domains     []string           `json:"domains" binding:"required"`
	CreatedAt   string             `json:"createdAt"`
	UpdatedAt   string             `json:"updatedAt"`
	Managers    []string           `json:"managers" binding:"required"`
}

// Delete deletes documents
func (mc *Application) Delete() (int64, error) {
	collection := client.Database(dbName).Collection("applications")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	deleteResult, err := collection.DeleteOne(ctx, bson.M{"_id": mc.ID})

	return deleteResult.DeletedCount, err
}

// Update updates documents
func (mc *Application) Update() (int64, error) {
	collection := client.Database(dbName).Collection("applications")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	u1 := uuid.Must(uuid.NewV4(), nil)
	mc.Secret = fmt.Sprintf("%s", u1)
	mc.UpdatedAt = time.Now().String()

	updateResult, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": mc.ID},
		bson.M{"$set": bson.M{"secret": mc.Secret}},
	)

	if err != nil {
		return 0, err
	}

	return updateResult.ModifiedCount, err
}

//Insert creates new document
func (mc *Application) Insert() (string, error) {
	collection := client.Database(dbName).Collection("applications")

	mc.ID = primitive.NewObjectID()
	mc.CreatedAt = time.Now().String()
	mc.UpdatedAt = ""
	u1 := uuid.Must(uuid.NewV4(), nil)
	mc.Secret = fmt.Sprintf("%s", u1)

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

	filter := find.ToBson()
	if len(mc.Name) > 0 {
		filter["name"] = mc.Name
	}
	if len(mc.Managers) > 0 {
		filter["managers"] = strings.Join(mc.Managers, ",")
	}

	total, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return result, 0, err
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return make([]interface{}, 0), 0, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		app := &Application{}
		err := cur.Decode(app)
		result = append(result, *app)
		if err != nil {
			return make([]interface{}, 0), 0, err
		}
	}

	return result, total, nil
}
