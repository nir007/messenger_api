package application

import (
	"context"
	"fmt"
	"time"

	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Manager struct {
	ID           string `json:"id"`
	Name         string `json:"name" binding:"required,max=50,min=1"`
	SecondName   string `json:"second" binding:"required,max=50,min=1"`
	Email        string `json:"email" binding:"required,email"`
	IsConfirmed  bool   `json:"isConfirmed"`
	ConfirmToken string `json:"-"`
	Phone        string `json:"phone"`
	Password     string `json:"password" binding:"required,min=6"`
	CreatedAt    string `json:"createdAt" binding:"-"`
	UpdatedAt    string `json:"UpdatedAt" binding:"-"`
}

func (m *Manager) Insert() (*mongo.InsertOneResult, error) {
	collection := client.Database("messenger").Collection("managers")
	m.CreatedAt = time.Now().String()
	m.UpdatedAt = ""
	hash, _ := bcrypt.GenerateFromPassword([]byte(m.Password), 14)
	m.Password = string(hash)
	u1 := uuid.Must(uuid.NewV4(), nil)
	m.ConfirmToken = fmt.Sprintf("%s", u1)

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, m)

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		m.ID = oid.String()
	}

	return res, err
}

func (m *Manager) FindOne(f map[string]interface{}) error {
	collection := client.Database("messenger").Collection("managers")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.M{}

	if val, ok := f["id"]; ok {
		filter["_id"] = val
	}
	if val, ok := f["email"]; ok {
		filter["email"] = val
	}

	fmt.Println("filter:", filter)

	err := collection.FindOne(ctx, filter).Decode(m)
	fmt.Println(m)

	return err
}
