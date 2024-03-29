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
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
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

	var employees []model.EmployeeReturn
	if err != nil {
		return paginatedPayload, err
	}
	for courser.Next(context.TODO()) {
		var employee model.EmployeeReturn
		err := courser.Decode(&employee)
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

func (c Client) GetEmployeesByDepartment(department string) ([]model.EmployeeReturn, error) {
	var employeeArr []model.EmployeeReturn
	filter := bson.M{"department": department}
	cur, err := c.Employee.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var employee model.EmployeeReturn
		err = cur.Decode(&employee)
		if err != nil {
			return nil, err
		}
		employeeArr = append(employeeArr, employee)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	return employeeArr, nil
}

func (c Client) UpdateEmpShift(update model.Shift, id string) (model.Employee, error) {
	filter := bson.M{"id": id}
	// datensatz zur id auslesen
	// check doc geschnitten datensatzen
	// change update

	employee := c.GetByID(id)
	newShifts := append(employee.Shifts, update)
	employee.Shifts = newShifts
	updater := bson.D{{"$set", employee}}

	results, err := c.Employee.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		log.Println("database error")
		return employee, err
	}
	if results.ModifiedCount == 0 {
		err = errors.New("No update could be send to the database")
		return employee, err
	}

	return employee, nil
}

func (c Client) UpdateEmp(update model.EmployeeReturn) (*mongo.UpdateResult, error) {
	filter := bson.M{"id": update.ID}
	// datensatz zur id auslesen
	// check doc geschnitten datensatzen
	// change update
	if update.ID == "" {
		IdMissing := fmt.Sprintf("User %v got no ID", update.ID)
		return nil, errors.New(IdMissing)
	}
	courser := c.Employee.FindOne(context.TODO(), filter)
	var employee model.Employee
	err := courser.Decode(&employee)
	if employee.ID == "" {
		IdWrong := fmt.Sprintf("User %v dosent exist", update.ID)
		return nil, errors.New(IdWrong)
	}
	fmt.Println(update)
	var setElements bson.D
	if update.FirstName != "" {
		fmt.Sprintf(update.FirstName)
		setElements = append(setElements, bson.E{Key: "firstname", Value: update.FirstName})
	}
	if update.LastName != "" {
		fmt.Sprintf(update.LastName)
		setElements = append(setElements, bson.E{Key: "lastname", Value: update.LastName})
	}
	if update.Email != "" {
		fmt.Sprintf(update.Email)
		setElements = append(setElements, bson.E{Key: "email", Value: update.Email})
	}
	setMap := bson.D{
		{"$set", setElements},
	}
	result, err := c.Employee.UpdateOne(context.TODO(), filter, setMap)
	if err != nil {
		return nil, err

	}
	return result, nil
}

/*updater := bson.D{{"$set", update}}

results, err := c.Employee.UpdateOne(context.TODO(), filter, updater)
if err != nil {
	log.Println("database error")
	return model.EmployeeReturn{}, err
}
if results.ModifiedCount == 0 {
	err = errors.New("No update could be send to the database")
	return model.EmployeeReturn{}, err
}

return update, nil*/
