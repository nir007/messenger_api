package dto

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoParamsGetter struct for getting params for searching
type MongoParamsGetter interface {
	ToBson() bson.M
	Limit() *int64
	Skip() *int64
}

// Pager for pagination
type Pager struct {
	Page    int64 `json:"page" form:"page" bson:"-"`
	PerPage int64 `json:"perpage" form:"perpage" bson:"-"`
}

// Limit returns number items on one page
func (f *Pager) Limit() *int64 {
	if f.PerPage == 0 || f.PerPage > 100 || f.PerPage < 0 {
		f.PerPage = 10
	}
	return &f.PerPage
}

// Skip returns number to skip documents
func (f *Pager) Skip() *int64 {
	skip := *f.Limit() * (f.Page - 1)

	if skip < 0 {
		skip = 0
	}

	return &skip
}

// FindApplications struct for searching applications
type FindApplications struct {
	Pager
	ID        primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
	Name      string             `json:"name" form:"name" bson:"name"`
	Managers  []string           `json:"managers" form:"managers" bson:"managers"`
	ManagerID string             `json:"managerId" form:"manager_id"`
}

// ToBson forms bson struct for searching documents
func (f *FindApplications) ToBson() bson.M {
	b, _ := bson.Marshal(f)
	var dataM bson.M
	_ = bson.Unmarshal(b, &dataM)

	return dataM
}

// FindUsers struct for finding users of application
type FindUsers struct {
	Pager
	ID            primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
	UID           string             `json:"uid" form:"uid" bson:"uid,omitempty"`
	Name          string             `json:"name" form:"name" bson:"name,omitempty"`
	Second        string             `json:"second" form:"second" bson:"second,omitempty"`
	ApplicationID string             `json:"applicationid" form:"applicationid" bson:"applicationid,omitempty"`
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
	Pager
	ID     primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
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
