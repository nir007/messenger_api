package application

import (
	"context"
	"messenger/dto"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Message struct
type Message struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ApplicationID string             `json:"applicationID" binding:"required"`
}

// Delete deletes documents
func (mc *Message) Delete() (int64, error) {
	collection := client.Database("messenger").Collection("messages_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	deleteResult, err := collection.DeleteOne(ctx, bson.M{"_id": mc.ID})

	return deleteResult.DeletedCount, err
}

// Update updates documents
func (mc *Message) Update() (int64, error) {
	return 0, nil
}

//Insert  creates new document
func (mc *Message) Insert() (string, error) {
	return "newId", nil
}

// FindOne finds one document
func (mc *Message) FindOne(find dto.MongoParamsGetter) error {
	return nil
}

// Find finds several documents
func (mc *Message) Find(find dto.MongoParamsGetter) ([]interface{}, int64, error) {
	return []interface{}{}, 0, nil
}
