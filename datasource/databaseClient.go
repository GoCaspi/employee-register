package datasource

import (
	"context"
	"errors"
	"example-project/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . MongoDBInterface
type MongoDBInterface interface {
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
}

type Client struct {
	Employee MongoDBInterface
}

func NewDbClient(d model.DbConfig) Client {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(d.URL))
	db := client.Database(d.Database)
	return Client{
		Employee: db.Collection("employee"),
	}
}

func (c Client) UpdateMany(docs []interface{}) (interface{}, error) {
	results, err := c.Employee.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Println("database error")
		return nil, err
	}
	return results.InsertedIDs, nil
}

func (c Client) GetByID(id string) model.Employee {
	filter := bson.M{"id": id}
	courser := c.Employee.FindOne(context.TODO(), filter)
	var employee model.Employee
	err := courser.Decode(&employee)
	if err != nil {
		log.Println("error during data marshalling")
	}
	return employee
}

func (c Client) DeleteByID(id string) (interface{}, error) {
	filter := bson.M{"id": id}

	results, err := c.Employee.DeleteOne(context.TODO(), filter)

	if err != nil {

		return nil, err
	}
	if results.DeletedCount == 0 {
		deleterror := errors.New("the Employee id is not existing")
		return nil, deleterror
	}
	return results.DeletedCount, nil
}

func (c Client) GetPaginated(page int, limit int) (model.PaginatedPayload, error) {
	var paginatedPayload model.PaginatedPayload
	skipMax, er := c.Employee.CountDocuments(context.TODO(), bson.D{})
	if er != nil {
		return model.PaginatedPayload{}, errors.New("error at counting documents")
	}

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"id", 1}})
	limit64 := int64(limit)
	var maxPages = float64(skipMax) / float64(limit64)
	maxPages = math.Ceil(maxPages)
	paginatedPayload.PageLimit = int(maxPages)
	if page == 0 || math.Signbit(float64(page)) {
		invalidPageNumber := errors.New("invalid page number, page number can't be zero or negative")
		return paginatedPayload, invalidPageNumber
	}
	if limit == 0 || math.Signbit(float64(limit)) {
		invalidPageNumber := errors.New("invalid limit, limit can't be zero or negative")
		return paginatedPayload, invalidPageNumber
	}
	if maxPages == 0 {
		formattedError := fmt.Sprintf("your page limit is too high. please reduce it to: %v", skipMax)
		return paginatedPayload, errors.New(formattedError)
	}
	if page > int(maxPages) {
		outOfRange := errors.New("page limit reached, please reduce the page number")
		return paginatedPayload, outOfRange
	}
	pageSet := (page - 1) * limit
	findOptions.SetLimit(limit64)
	findOptions.SetSkip(int64(pageSet))
	courser, err := c.Employee.Find(context.TODO(), bson.D{}, findOptions)

	var employees []model.Employee
	if err != nil {
		return paginatedPayload, err
	}
	for courser.Next(context.TODO()) {
		var employee model.Employee
		err = courser.Decode(&employee)
		if err != nil {
			return paginatedPayload, err
		}
		employees = append(employees, employee)
	}
	if len(employees) == 0 {
		noEmployeesError := errors.New("no employees exist")
		return paginatedPayload, noEmployeesError
	}
	paginatedPayload.Employees = employees
	return paginatedPayload, nil

}
