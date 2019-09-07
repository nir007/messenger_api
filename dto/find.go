package dto

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindApplications struct for searching applications
type FindApplications struct {
	MyBSON    `bson:"-"`
	Page      `bson:"-"`
	ID        primitive.ObjectID `json:"id" form:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" form:"name" bson:"name,omitempty"`
	Secret    string             `json:"secret" form:"secret" bson:"secret,omitempty"`
	Managers  []string           `json:"managers" form:"managers" bson:"managers,omitempty"`
	ManagerID string             `json:"managerID" form:"managerid" bson:"-"`
	Domains   []string           `json:"domains" bson:"domains,omitempty"`
}

// ToBson forms bson struct for searching documents
func (f *FindApplications) ToBson() bson.M {
	/*if len(mc.Managers) > 0 {
		filter["managers"] = strings.Join(mc.Managers, ",")
	}*/

	b, _ := bson.Marshal(f)
	var dataM bson.M
	_ = bson.Unmarshal(b, &dataM)

	return dataM
}

// FindUsers struct for finding users of application
type FindUsers struct {
	MyBSON        `bson:"-"`
	Page          `bson:"-"`
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UID           string             `json:"uid" form:"uid" bson:"uid,omitempty"`
	Name          string             `json:"name" form:"name" bson:"name,omitempty"`
	Second        string             `json:"second" form:"second" bson:"second,omitempty"`
	ApplicationID string             `json:"applicationId" form:"applicationid" bson:"applicationid,omitempty"`
	BlackList     []string           `json:"blackList"`
	Email         string             `json:"email" form:"email" bson:"email,omitempty"`
	Phone         string             `json:"phone" form:"phone" bson:"phone,omitempty"`
}

// ToBson forms bson struct for searching documents
func (f *FindUsers) ToBson() bson.M {
	b, _ := bson.Marshal(f)
	var dataM bson.M
	_ = bson.Unmarshal(b, &dataM)

	return dataM
}

// FindManagers struct for finding users of application
type FindManagers struct {
	MyBSON `bson:"-"`
	Page   `bson:"-"`
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" form:"name" bson:"name,omitempty"`
	Second string             `json:"second" form:"second" bson:"second,omitempty"`
	Email  string             `json:"email" form:"email" bson:"email,omitempty"`
	Phone  string             `json:"phone" form:"phone" bson:"phone,omitempty"`
}

// ToBson forms bson struct for searching documents
func (f *FindManagers) ToBson() bson.M {
	b, _ := bson.Marshal(f)
	var dataM bson.M
	_ = bson.Unmarshal(b, &dataM)

	return dataM
}

// FindDialogs struct for finding dialogs of application
type FindDialogs struct {
	MyBSON        `bson:"-"`
	Page          `bson:"-"`
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	LastMessage   string             `json:"lastMessage" form:"lastMessage" bson:"lastmessage,omitempty"`
	ApplicationID string             `json:"applicationID" form:"applicationID" bson:"applicationid,omitempty"`
	UID1          string             `json:"uid1" form:"uid1" bson:"uid1,omitempty"`
	UID2          string             `json:"uid2" form:"uid2" bson:"uid2,omitempty"`
	IsRed         bool               `json:"isRed" form:"isRed" bson:"isred,omitempty"`
}

// ToBson forms bson struct for searching documents
func (f *FindDialogs) ToBson() bson.M {
	b, _ := bson.Marshal(f)
	var dataM bson.M
	_ = bson.Unmarshal(b, &dataM)

	uid1, ok1 := dataM["uid1"]
	uid2, ok2 := dataM["uid2"]

	if ok1 && ok2 {
		delete(dataM, "uid1")
		delete(dataM, "uid2")
		dataM["$or"] = []bson.M{
			bson.M{
				"uid1": uid1,
				"uid2": uid2,
			},
			bson.M{
				"uid1": uid2,
				"uid2": uid1,
			},
		}
	}
	return dataM
}

// FindMessages struct for searching messages of dialog
type FindMessages struct {
	MyBSON        `bson:"-"`
	Page          `bson:"-"`
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Text          string             `json:"text" form:"text" bson:"text,omitempty"`
	DialogID      string             `json:"dialogID" form:"dialogID" bson:"dialogid,omitempty"`
	ApplicationID string             `json:"applicationID" form:"applicationID" bson:"applicationid,omitempty"`
	UID1          string             `json:"uid1" form:"uid1" bson:"uid1,omitempty"`
	UID2          string             `json:"uid2" form:"uid2" bson:"uid2,omitempty"`
	IsRed         bool               `json:"isRed" form:"isRed" bson:"isred,omitempty"`
}

// ToBson forms bson struct for searching documents
func (f *FindMessages) ToBson() bson.M {
	b, _ := bson.Marshal(f)
	var dataM bson.M
	_ = bson.Unmarshal(b, &dataM)

	uid1, ok1 := dataM["uid1"]
	uid2, ok2 := dataM["uid2"]

	if ok1 && ok2 {
		delete(dataM, "uid1")
		delete(dataM, "uid2")
		dataM["$or"] = []bson.M{
			bson.M{
				"uid1": uid1,
				"uid2": uid2,
			},
			bson.M{
				"uid1": uid2,
				"uid2": uid1,
			},
		}
	}
	return dataM
}
