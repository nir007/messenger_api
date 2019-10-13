package drepository

import (
	"context"
	"errors"
	"fmt"
	"messenger/dto"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MessageFile struct
type MessageFile struct {
	URL        string `json:"url"`
	PreviewURL string `json:"previewUrl"`
	Type       string `json:"type"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Weight     int    `json:"weight"`
}

// Message struct
type Message struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	DialogID      string             `json:"dialogID"`
	UID1          string             `json:"uid1" binding:"required"`
	UID2          string             `json:"uid2" binding:"required"`
	Text          string             `json:"text" binding:"required"`
	Files         []MessageFile      `json:"files" binding:"-"`
	IsRed         bool               `json:"isRed"`
	ApplicationID string             `json:"applicationId"`
	CreatedAt     string             `json:"createdAt" binding:"-"`
	UpdatedAt     string             `json:"updatedAt" binding:"-"`
	DeletedAt     string             `json:"deletedAt" binding:"-"`
}

// Delete deletes documents
func (mc *Message) Delete() (int64, error) {
	collection := client.Database(dbName).Collection("messages_" + mc.ApplicationID)
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

	if _, err = dialog.Update(findDialog, updateDialog); err != nil {
		_, err = dialog.Insert()
		if err != nil {
			return "", err
		}
	}

	collection := client.Database(dbName).Collection("messages_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	mc.ID = primitive.NewObjectID()
	mc.DialogID = dialog.ID.Hex()
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
	result := make([]interface{}, 0)

	fmt.Println("find.ToBson(): ", find.ToBson())

	collection := client.Database(dbName).Collection("messages_" + mc.ApplicationID)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	options := &options.FindOptions{Skip: find.Skip(), Limit: find.Limit(), Sort: find.Sort()}

	total, err := collection.CountDocuments(ctx, find.ToBson())
	if err != nil {
		return result, 0, err
	}

	cur, err := collection.Find(ctx, find.ToBson(), options)
	if err != nil {
		return make([]interface{}, 0), 0, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		app := &Message{}
		err := cur.Decode(app)
		result = append(result, *app)
		if err != nil {
			return make([]interface{}, 0), 0, err
		}
	}

	return result, total, nil
}
