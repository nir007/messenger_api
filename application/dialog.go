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
	CreatedAt     string             `json:"createdAt" binding:"-"`
	UpdatedAt     string             `json:"updatedAt" binding:"-"`
	ApplicationID string             `json:"applicationID" binding:"required"`
	IsRed         bool               `json:"isRed" binding:"required"`
}

// Delete deletes documents
func (mc *Dialog) Delete() (int64, error) {
	collection := client.Database("messenger").Collection("dialogs_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	deleteResult, err := collection.DeleteOne(ctx, bson.M{"_id": mc.ID})

	return deleteResult.DeletedCount, err
}

// Update updates documents
func (mc *Dialog) Update() (int64, error) {
	return 0, nil
}

//Insert creates new document
func (mc *Dialog) Insert() (string, error) {
	find := &dto.FindUsers{ID: mc.ID}
	if err := mc.FindOne(find); err == nil {
		updatedRows, err := mc.updateLastMessage()

		if err != nil || updatedRows == 0 {
			return "", errors.New("updated rows: 0, " + err.Error())
		}

		return mc.ID.Hex(), nil
	}

	collection := client.Database("messenger").Collection("dialogs_" + mc.ApplicationID)

	mc.ID = primitive.NewObjectID()
	mc.CreatedAt = time.Now().String()
	mc.UpdatedAt = ""

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, mc)

	return res.InsertedID.(string), err
}

// FindOne finds one document
func (mc *Dialog) FindOne(find dto.MongoParamsGetter) error {
	collection := client.Database("messenger").Collection("dialogs_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	err := collection.FindOne(ctx, find).Decode(mc)

	return err
}

// Find finds several documents by pages
func (mc *Dialog) Find(find dto.MongoParamsGetter) ([]interface{}, int64, error) {
	return []interface{}{}, 0, nil
}

func (mc *Dialog) updateLastMessage() (int64, error) {
	collection := client.Database("messenger").Collection("dialogs_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	updateResult, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": mc.ID},
		bson.M{"$set": bson.M{"lastmessage": mc.LastMessage}},
	)

	if err != nil {
		return 0, err
	}

	return updateResult.ModifiedCount, err
}
