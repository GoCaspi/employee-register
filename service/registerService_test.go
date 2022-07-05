package service_test

import (
	"errors"
	"example-project/model"
	"example-project/service"
	"example-project/service/servicefakes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEmployeeById(t *testing.T) {

	fakeDB := &servicefakes.FakeDatabaseInterface{}
	fakeErr := errors.New("")
	data := model.Employee{
		ID:        "1",
		FirstName: "jon",
		LastName:  "kock",
		Email:     "jon@gmail.com",
	}

	var tests = []struct {
		hasError  bool
		Data      model.Employee
		FakeError error
	}{
		{false, data, nil},
		{true, data, fakeErr},
	}

	for _, tt := range tests {

		if tt.hasError {
			fakeDB.GetByIDReturns(tt.Data, tt.FakeError)
			serviceInstance := service.NewEmployeeService(fakeDB)
			actual, err := serviceInstance.GetEmployeeById(tt.Data.ID)
			assert.Equal(t, tt.Data, actual)
			assert.Equal(t, tt.FakeError, err)
		}

		fakeDB.GetByIDReturns(tt.Data, tt.FakeError)
		serviceInstance := service.NewEmployeeService(fakeDB)
		actual, err := serviceInstance.GetEmployeeById(tt.Data.ID)
		assert.Equal(t, data, actual)
		assert.Equal(t, tt.FakeError, err)
	}

}

func TestCreateEmployees(t *testing.T) {
	//here comes your first unit test which should cover the function CreateEmployees

	fakeDB := &servicefakes.FakeDatabaseInterface{}
	fakeEmployees := []model.Employee{
		model.Employee{ID: "1", FirstName: "jon", LastName: "doe", Email: "jondoe@mail.com"},
	}
	fakeErr := errors.New("")
	fakeInterface := struct {
		InsertedIDs int
	}{
		1,
	}

	var tests = []struct {
		Fakes     []model.Employee
		Err       error
		Result    struct{ InsertedIDs int }
		ReturnErr bool
	}{
		{fakeEmployees, nil, fakeInterface, false},
		{fakeEmployees, fakeErr, fakeInterface, true},
	}

	for _, tt := range tests {
		if !tt.ReturnErr {
			fakeDB.UpdateManyReturns(tt.Result, tt.Err)

			serviceInstance := service.NewEmployeeService(fakeDB)
			actual, _ := serviceInstance.CreateEmployees(tt.Fakes)
			assert.Equal(t, tt.Result, actual)
		}

		if tt.ReturnErr {
			fakeDB.UpdateManyReturns(tt.Result, tt.Err)

			serviceInstance := service.NewEmployeeService(fakeDB)
			actual, actualErr := serviceInstance.CreateEmployees(tt.Fakes)
			assert.Equal(t, nil, actual)
			assert.Equal(t, tt.Err, actualErr)
		}
	}

}
