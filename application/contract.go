package application

import "messenger/dto"

// MongoDbManager for application structs
type MongoDbManager interface {
	// Delete deletes documents
	Delete() (int64, error)
	// Update updates documents
	Update() (int64, error)
	//Insert  creates new document
	Insert() (string, error)
	// FindOne finds one document
	FindOne(find dto.MongoParamsGetter) error
	// Find finds several documents by pages
	Find(find dto.MongoParamsGetter) ([]interface{}, int64, error)
}
