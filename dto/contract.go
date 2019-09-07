package dto

import "go.mongodb.org/mongo-driver/bson"

// SearchParamsGetter struct for getting params for searching
type SearchParamsGetter interface {
	BSONMaker
	Pager
}

// BSONMaker struct for getting params for searching
type BSONMaker interface {
	ToBson() bson.M
}

// MyBSON struct for making bson
type MyBSON struct{}

// ToBson forms bson struct for searching documents
func (f *MyBSON) ToBson() bson.M {
	b, _ := bson.Marshal(f)
	var dataM bson.M
	_ = bson.Unmarshal(b, &dataM)

	return dataM
}

// Pager for pagination
type Pager interface {
	Limit() *int64
	Skip() *int64
	Sort() bson.M
}

// Page for pagination
type Page struct {
	Page      int64  `json:"page" form:"page" bson:"-"`
	PerPage   int64  `json:"perpage" form:"perpage" bson:"-"`
	SortField string `json:"sort" form:"sort" bson:"-"`
	SortOrder string `json:"order" form:"order" bson:"-"`
}

// Limit returns number items on one page
func (f *Page) Limit() *int64 {
	if f.PerPage == 0 || f.PerPage > 100 || f.PerPage < 0 {
		f.PerPage = 10
	}
	return &f.PerPage
}

// Skip returns number to skip documents
func (f *Page) Skip() *int64 {
	skip := *f.Limit() * (f.Page - 1)

	if skip < 0 {
		skip = 0
	}

	return &skip
}

// Sort return options for sorting
func (f *Page) Sort() bson.M {
	if len(f.SortField) == 0 || (f.SortOrder != "asc" && f.SortOrder != "desc") {
		return bson.M{"createdat": -1}
	}

	orders := map[string]int{
		"asc":  1,
		"desc": -1,
	}

	return bson.M{f.SortField: orders[f.SortOrder]}
}
