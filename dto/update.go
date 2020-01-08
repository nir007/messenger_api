package dto

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateApplication struct for for updating application
type UpdateApplication struct {
	ID          primitive.ObjectID `json:"id" binding:"-" bson:"_id,omitempty"`
	Name        string             `json:"name" binding:"max=50,min=1" bson:"name,omitempty"`
	Description string             `json:"description" binding:"max=100,min=1" bson:"description,omitempty"`
	Domains     []string           `json:"domains" bson:"domains,omitempty"`
	IsActive    bool               `json:"isActive" bson:"isactive,omitempty"`
	Managers    []string           `json:"managers" bson:"managers,omitempty"`
	Secret      string             `json:"secret" binding:"-" bson:"secret,omitempty"`
	Salt        string             `json:"salt" bson:"salt,omitempty"`
	UpdatedAt   string             `json:"updatedAt" binding:"-" bson:"updatedat,omitempty"`
}

// ToBson forms bson struct for searching documents
func (f *UpdateApplication) ToBson() bson.M {
	f.UpdatedAt = time.Now().String()
	b, _ := bson.Marshal(f)

	fmt.Println(f)

	var dataM bson.M
	_ = bson.Unmarshal(b, &dataM)

	if !f.IsActive {
		dataM["isactive"] = false
	}

	if len(f.Domains) == 0 {
		dataM["domains"] = []string{}
	}

	fmt.Println(dataM)
	return dataM
}

// UpdateUser struct for updating application user
type UpdateUser struct {
	ID         primitive.ObjectID `json:"id" binding:"-" bson:"_id,omitempty"`
	UID        string             `json:"uid" bson:"uid,omitempty"`
	Name       string             `json:"name" binding:"max=50,min=1" bson:"name,omitempty"`
	SecondName string             `json:"second" binding:"max=100,min=1" bson:"description,omitempty"`
	BlackList  []string           `json:"blackList" bson:"blacklist,omitempty"`
	Email      string             `json:"email" bson:"email,omitempty"`
	Phone      string             `json:"phone"  bson:"phone,omitempty"`
	UpdatedAt  string             `json:"updatedAt" binding:"-" bson:"updatedat,omitempty"`
}

// ToBson forms bson struct for searching documents
func (f *UpdateUser) ToBson() bson.M {
	f.UpdatedAt = time.Now().String()
	b, _ := bson.Marshal(f)

	var dataM bson.M
	_ = bson.Unmarshal(b, &dataM)

	return dataM
}

// UpdateDialog struct for for updating dialog
type UpdateDialog struct {
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
	Name        string             `json:"name" binding:"max=50,min=1"`
	SecondName  string             `json:"second" binding:"max=50,min=1"`
	Email       string             `json:"email" binding:"email"`
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
