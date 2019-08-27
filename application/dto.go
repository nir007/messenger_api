package application

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type MongoParamsGetter interface {
	ToBson() bson.M
	Limit() *int64
	Skip() *int64
}

type UpdateApplicationSecret struct {
	ID primitive.ObjectID `json:"id" binding:"required"`
}

type FindApplications struct {
	Name     string   `json:"name"`
	Managers []string `json:"managers"`
}

func (f *FindApplications) ToBson() bson.M {
	return bson.M{}
}

type FindUsers struct {
	Name          string `json:"name" form:"name" bson:"name,omitempty"`
	Second        string `json:"second" form:"second" bson:"second,omitempty"`
	ApplicationID string `json:"applicationid" form:"applicationid" bson:"applicationid,omitempty"`
	Page          int64  `json:"page" form:"page" bson:"-"`
	PerPage       int64  `json:"perpage" form:"perpage" bson:"-"`
	Email         string `json:"email" form:"email" bson:"email,omitempty"`
	Phone         string `json:"phone" form:"phone" bson:"phone,omitempty"`
}

func (f *FindUsers) ToBson() bson.M {
	b, _ := bson.Marshal(f)
	var dataM bson.M
	_ = bson.Unmarshal(b, &dataM)

	return dataM
}

func (f *FindUsers) Limit() *int64 {
	if f.PerPage == 0 || f.PerPage > 100 || f.PerPage < 0 {
		f.PerPage = 10
	}
	return &f.PerPage
}

func (f *FindUsers) Skip() *int64 {
	skip := *f.Limit() * (f.Page - 1)

	if skip < 0 {
		skip = 0
	}

	return &skip
}
