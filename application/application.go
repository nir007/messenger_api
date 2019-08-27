package application

import (
	"context"
	"fmt"
	"time"

	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type Application struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Secret      string    `json:"secret"`
	Domains     []string  `json:"domains" binding:"required"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Managers    []string `json:"managers" binding:"required"`
}

func (a *Application) Insert(c *gin.Context) (*mongo.InsertOneResult, error) {
	collection := client.Database("messenger").Collection("applications")

	a.ID = primitive.NewObjectID()
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	u1 := uuid.Must(uuid.NewV4(), nil)
	a.Secret = fmt.Sprintf("%s", u1)

	if managerId, ok := c.Get("managerId"); !ok {
		return nil, fmt.Errorf("%s", "undefined manager id")
	} else {
		a.Managers = append(a.Managers, managerId.(string))
		ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		res, err := collection.InsertOne(ctx, a)

		return res, err
	}
}

func (a *Application) Find(c *gin.Context) ([]Application, error) {
	result := make([]Application, 0)

	collection := client.Database("messenger").Collection("applications")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.M{}
	if len(a.Name) > 0 {
		filter["name"] = a.Name
	}
	if len(a.Managers) > 0 {
		filter["managers"] = strings.Join(a.Managers, ",")
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil { log.Fatal(err) }

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		app := &Application{}
		err := cur.Decode(app)
		result = append(result, *app)
		if err != nil {
			log.Fatal(err)
		}
	}

	return result, err
}

func (a *Application) FindOne(c *gin.Context) error {
	collection := client.Database("messenger").Collection("applications")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	return collection.FindOne(ctx, bson.M{"_id": a.ID}).Decode(a)
}

func (a *Application) Update(c *gin.Context) (int64, error) {
	collection := client.Database("messenger").Collection("applications")
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	u1 := uuid.Must(uuid.NewV4(), nil)
	a.Secret = fmt.Sprintf("%s", u1)

	updateResult, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": a.ID},
		bson.M{"$set": bson.M{"secret": a.Secret}},
	)

	if err != nil {
		return 0, err
	}

	return updateResult.ModifiedCount, err
}
