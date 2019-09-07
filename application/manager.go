package application

import (
	"context"
	"errors"
	"fmt"
	"messenger/dto"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// Manager struct for apps managers
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
	UpdatedAt    string             `json:"updatedAt" binding:"-"`
}

// Delete deletes documents
func (mc *Manager) Delete() (int64, error) {
	collection := client.Database(dbName).Collection("managers")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	deleteResult, err := collection.DeleteOne(ctx, bson.M{"_id": mc.ID})

	return deleteResult.DeletedCount, err
}

//Update changes document
func (mc *Manager) Update(find dto.SearchParamsGetter, update dto.BSONMaker) (int64, error) {
	return 0, nil
}

//Insert creates new document
func (mc *Manager) Insert() (string, error) {
	find := &dto.FindManagers{Email: mc.Email}
	if err := mc.FindOne(find); err == nil {
		return "", errors.New("user already exists")
	}

	collection := client.Database(dbName).Collection("managers")

	mc.ID = primitive.NewObjectID()
	mc.CreatedAt = time.Now().String()
	mc.UpdatedAt = ""
	hash, _ := bcrypt.GenerateFromPassword([]byte(mc.Password), 14)
	mc.Password = string(hash)

	u1 := uuid.Must(uuid.NewV4(), nil)
	mc.ConfirmToken = fmt.Sprintf("%s", u1)

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, mc)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", res.InsertedID), err
}

// FindOne finds one document
func (mc *Manager) FindOne(find dto.SearchParamsGetter) error {
	collection := client.Database(dbName).Collection("managers")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	fmt.Println(find.ToBson())

	return collection.FindOne(ctx, find.ToBson()).Decode(mc)
}

// Find finds several documents by pages
func (mc *Manager) Find(find dto.SearchParamsGetter) ([]interface{}, int64, error) {
	result := make([]interface{}, 0)

	collection := client.Database(dbName).Collection("managers")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	options := &options.FindOptions{Skip: find.Skip(), Limit: find.Limit(), Sort: find.Sort()}
	total, err := collection.CountDocuments(ctx, find.ToBson())
	if err != nil {
		return result, 0, err
	}

	cur, err := collection.Find(ctx, find.ToBson(), options)
	if err != nil {
		return make([]interface{}, 0), 0, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		app := &Manager{}
		err := cur.Decode(app)
		result = append(result, *app)
		if err != nil {
			return make([]interface{}, 0), 0, err
		}
	}

	return result, total, nil
}
