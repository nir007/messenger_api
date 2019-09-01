package application

import "messenger/dto"

// MongoDbManager for application structs
type MongoDbManager interface {
	// Delete deletes documents
	Delete() (int64, error)
	// Update updates documents
	Update(find dto.SearchParamsGetter, update dto.BSONMaker) (int64, error)
	//Insert  creates new document
	Insert() (string, error)
	// FindOne finds one document
	FindOne(find dto.SearchParamsGetter) error
	// Find finds several documents by pages
	Find(find dto.SearchParamsGetter) ([]interface{}, int64, error)
}
