package datasource

import (
	"context"
	"errors"
	"example-project/model"
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

func (c Client) GetAllPaginated(pagenumber int, limit int) ([]model.Employee, error) {

	// var counter is increased by one for every document that was found in the database. It is used to calculate the last page.
	var counter int = 0
	indicatorFilter := bson.M{}
	indicator, _ := c.Employee.Find(context.TODO(), indicatorFilter)
	defer indicator.Close(context.Background())
	for indicator.Next(context.Background()) {
		counter++
	}
	lastPageC := math.Ceil(float64(counter / limit))
	lastPageF := math.Floor(float64(counter / limit))
	lastPage := lastPageC
	if lastPageC == lastPageF {
		lastPage++
	}

	if float64(pagenumber) > lastPage {
		outOfBoundsErr := errors.New("The given pagenumber exceeds the total number of pages avaible.")
		log.Println(outOfBoundsErr)
		return nil, outOfBoundsErr
	}
	// setting the findoptions to the pagination structure given as parameters
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"id", 1}})
	limit64 := int64(limit)
	findOptions.SetLimit(limit64)
	findOptions.SetSkip(int64((int64(pagenumber) - 1) * limit64))
	var employeeArr []model.Employee
	filter := bson.M{}
	cur, err := c.Employee.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var employee model.Employee
		err := cur.Decode(&employee)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		employeeArr = append(employeeArr, employee)
	}
	if err := cur.Err(); err != nil {
		log.Print(err)
		return nil, err
	}

	return employeeArr, nil
}
