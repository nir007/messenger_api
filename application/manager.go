package application

import (
	"context"
	"fmt"
	"time"

	"errors"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Manager struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name" binding:"required,max=50,min=1"`
	SecondName   string             `json:"second" binding:"required,max=50,min=1"`
	Email        string             `json:"email" binding:"required,email"`
	IsConfirmed  bool               `json:"isConfirmed" binding:"-"`
	ConfirmToken string             `json:"-" binding:"-"`
	Phone        string             `json:"phone"`
	Password     string             `json:"password" binding:"required,min=6"`
	CreatedAt    string             `json:"createdAt" binding:"-"`
	UpdatedAt    string             `json:"UpdatedAt" binding:"-"`
}

func (m *Manager) Insert() (*mongo.InsertOneResult, error) {
	if err := m.FindOne(bson.M{"email": m.Email}); err == nil {
		return nil, errors.New("user already exists")
	}

	collection := client.Database("messenger").Collection("managers")

	m.ID = primitive.NewObjectID()
	m.CreatedAt = time.Now().String()
	m.UpdatedAt = ""
	hash, _ := bcrypt.GenerateFromPassword([]byte(m.Password), 14)
	m.Password = string(hash)

	u1 := uuid.Must(uuid.NewV4(), nil)
	m.ConfirmToken = fmt.Sprintf("%s", u1)

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, m)

	return res, err
}

func (m *Manager) FindOne(find bson.M) error {
	collection := client.Database("messenger").Collection("managers")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	err := collection.FindOne(ctx, find).Decode(m)

	return err
}
