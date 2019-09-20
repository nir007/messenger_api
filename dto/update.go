package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateApplicationSecret struct for for updating secret key of application
type UpdateApplicationSecret struct {
	ID primitive.ObjectID `json:"id" binding:"required"`
}

// UpdateDialog struct for for updating dialog
type UpdateDialog struct {
	MyBSON      `bson:"-"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UID1        string             `json:"uid1" bson:"uid1,omitempty"`
	UID2        string             `json:"uid2" bson:"uid2,omitempty"`
	LastMessage string             `json:"lastMessage" bson:"lastmessage,omitempty"`
	IsRed       bool               `json:"isRed" bson:"isred,omitempty"`
	UpdatedAt   string             `json:"updatedAt" binding:"-" bson:"updatedat,omitempty"`
}

// ToBson forms bson struct for searching documents
func (f *UpdateDialog) ToBson() bson.M {
	f.UpdatedAt = time.Now().String()
	b, _ := bson.Marshal(f)
	var dataM bson.M
	_ = bson.Unmarshal(b, &dataM)

	return dataM
}

// UpdateManager struct for updatind manager data
type UpdateManager struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" binding:"required,max=50,min=1"`
	SecondName  string             `json:"second" binding:"required,max=50,min=1"`
	Email       string             `json:"email" binding:"required,email"`
	IsConfirmed bool               `json:"isConfirmed" binding:"-"`
	Phone       string             `json:"phone"`
	UpdatedAt   string             `json:"updatedAt" binding:"-" bson:"updatedat,omitempty"`
}

// ToBson forms bson struct for searching documents
func (f *UpdateManager) ToBson() bson.M {
	f.UpdatedAt = time.Now().String()
	b, _ := bson.Marshal(f)
	var dataM bson.M
	_ = bson.Unmarshal(b, &dataM)

	return dataM
}
