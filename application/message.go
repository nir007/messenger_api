package application

import (
	"context"
	"errors"
	"messenger/dto"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Message struct
type Message struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DialogID      string             `json:"dialogID"`
	UID1          string             `json:"uid1" binding:"required"`
	UID2          string             `json:"uid2" binding:"required"`
	Text          string             `json:"text" binding:"required"`
	IsRed         bool               `json:"isRed"`
	ApplicationID string             `json:"applicationID" binding:"required"`
	CreatedAt     string             `json:"createdAt" binding:"-"`
	UpdatedAt     string             `json:"updatedAt" binding:"-"`
}

// Delete deletes documents
func (mc *Message) Delete() (int64, error) {
	collection := client.Database(dbName).Collection("messages_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	deleteResult, err := collection.DeleteOne(ctx, bson.M{"_id": mc.ID})

	return deleteResult.DeletedCount, err
}

// Update updates documents
func (mc *Message) Update(find dto.SearchParamsGetter, update dto.BSONMaker) (int64, error) {
	return 0, nil
}

//Insert  creates new document
func (mc *Message) Insert() (string, error) {
	if len(mc.UID1) == 0 || len(mc.UID2) == 0 {
		return "", errors.New("uid1 and uid2 is required")
	}

	if mc.UID1 == mc.UID2 {
		return "", errors.New("uid1 equals uid2")
	}

	dialog := &Dialog{
		ApplicationID: mc.ApplicationID,
		UID1:          mc.UID1,
		UID2:          mc.UID2,
		LastMessage:   mc.Text,
	}

	findDialog := &dto.FindDialogs{
		ApplicationID: mc.ApplicationID,
		UID1:          mc.UID1,
		UID2:          mc.UID2,
	}

	updateDialog := &dto.UpdateDialog{
		UID1:        mc.UID1,
		UID2:        mc.UID2,
		LastMessage: mc.Text,
		UpdatedAt:   time.Now().String(),
	}

	var dialogID string

	if _, err := dialog.Update(findDialog, updateDialog); err != nil {
		dialogID, err = dialog.Insert()
		if err != nil {
			return "", err
		}
	}

	collection := client.Database(dbName).Collection("messages_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	mc.ID = primitive.NewObjectID()
	mc.DialogID = dialogID
	mc.CreatedAt = time.Now().String()
	mc.UpdatedAt = ""

	_, err := collection.InsertOne(ctx, mc)

	if err != nil {
		return "", err
	}

	return mc.ID.Hex(), nil
}

// FindOne finds one document
func (mc *Message) FindOne(find dto.SearchParamsGetter) error {
	return nil
}

// Find finds several documents
func (mc *Message) Find(find dto.SearchParamsGetter) ([]interface{}, int64, error) {
	return []interface{}{}, 0, nil
}
