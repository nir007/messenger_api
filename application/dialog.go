package application

import (
	"context"
	"errors"
	"messenger/dto"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Dialog struct for dialogs
type Dialog struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UID1          string             `json:"uid1" binding:"required"`
	UID2          string             `json:"uid2" binding:"required"`
	LastMessage   string             `json:"lastMessage" binding:"required"`
	ApplicationID string             `json:"applicationID" binding:"required"`
	IsRed         bool               `json:"isRed" binding:"required"`
	CreatedAt     string             `json:"createdAt" binding:"-"`
	UpdatedAt     string             `json:"updatedAt" binding:"-"`
}

// Delete deletes documents
func (mc *Dialog) Delete() (int64, error) {
	collection := client.Database(dbName).Collection("dialogs_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	deleteResult, err := collection.DeleteOne(ctx, bson.M{"_id": mc.ID})

	return deleteResult.DeletedCount, err
}

// Update updates documents
func (mc *Dialog) Update(find dto.SearchParamsGetter, update dto.BSONMaker) (int64, error) {
	collection := client.Database(dbName).Collection("dialogs_" + mc.ApplicationID)
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
		return 0, errors.New("undefined doalog")
	}

	_ = mc.FindOne(find)
	return updateResult.ModifiedCount, nil
}

//Insert creates new document
func (mc *Dialog) Insert() (string, error) {
	collection := client.Database(dbName).Collection("dialogs_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	mc.ID = primitive.NewObjectID()
	mc.CreatedAt = time.Now().String()
	mc.UpdatedAt = ""
	_, err := collection.InsertOne(ctx, mc)

	return mc.ID.Hex(), err
}

// FindOne finds one document
func (mc *Dialog) FindOne(find dto.SearchParamsGetter) error {
	collection := client.Database(dbName).Collection("dialogs_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	return collection.FindOne(ctx, find).Decode(mc)
}

// Find finds several documents by pages
func (mc *Dialog) Find(find dto.SearchParamsGetter) ([]interface{}, int64, error) {
	return []interface{}{}, 0, nil
}
