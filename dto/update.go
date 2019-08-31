package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateApplicationSecret struct for for updating secret key of application
type UpdateApplicationSecret struct {
	ID primitive.ObjectID `json:"id" binding:"required"`
}
