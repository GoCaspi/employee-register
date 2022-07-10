package service_test

import (
	"errors"
	"example-project/model"
	"example-project/service"
	"example-project/service/servicefakes"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestGetEmployeeById(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}

	data := model.Employee{
		ID:        "1",
		FirstName: "jon",
		LastName:  "kock",
		Email:     "jon@gmail.com",
	}

	fakeDB.GetByIDReturns(data)

	serviceInstance := service.NewEmployeeService(fakeDB)
	actual := serviceInstance.GetEmployeeById("1")
	assert.Equal(t, data, actual)

}

func TestCreateEmployees_Returns_valid_StatusCode(t *testing.T) {
	mockEmployees := []model.Employee{
		model.Employee{ID: "1", FirstName: "Joe", LastName: "Schmo", Email: "Joeschmo@mail.com", Auth: model.HashedAuth{}},
	}
	fakeDB := &servicefakes.FakeDatabaseInterface{}
	fakeDB.UpdateManyReturns(&mongo.InsertManyResult{}, nil)
	serviceInstance := service.NewEmployeeService(fakeDB)
	actual, _ := serviceInstance.CreateEmployees(mockEmployees)
	assert.NotNil(t, actual)
}

func TestCreateEmployees_Returns_Invalid_StatusCode400(t *testing.T) {
	mockEmployees := []model.Employee{
		model.Employee{ID: "1", FirstName: "Joe", LastName: "Schmo", Email: "Joeschmo@mail.com", Auth: model.HashedAuth{}},
	}
	fakeDB := &servicefakes.FakeDatabaseInterface{}
	fakeDB.UpdateManyReturns(&mongo.InsertManyResult{}, errors.New(""))
	serviceInstance := service.NewEmployeeService(fakeDB)
	_, actualError := serviceInstance.CreateEmployees(mockEmployees)
	assert.NotNil(t, actualError)
}

func TestDeleteEmployee(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}
	fakeDB.DeleteByIDReturns(&mongo.DeleteResult{DeletedCount: 1}, nil)

	serviceInstance := service.NewEmployeeService(fakeDB)
	actual, _ := serviceInstance.DeleteEmployee("1")
	assert.NotNil(t, actual)
}
