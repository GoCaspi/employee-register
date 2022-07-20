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

func TestGetPaginatedEmployees(t *testing.T) {
	fakeDB := &servicefakes.FakeDatabaseInterface{}
	fakePaginatedPayload := model.PaginatedPayload{
		PageLimit: 2,
		Employees: []model.EmployeeReturn{
			model.EmployeeReturn{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
			model.EmployeeReturn{ID: "200", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
		},
	}
	fakeDB.GetPaginatedReturns(fakePaginatedPayload, nil)
	serviceInstance := service.NewEmployeeService(fakeDB)
	actual, err := serviceInstance.GetPaginatedEmployees(1, 2)
	assert.Equal(t, fakePaginatedPayload, actual, err)
}

func TestEmployeeService_GetEmployeesDepartmentFilter(t *testing.T) {
	fakeDb := &servicefakes.FakeDatabaseInterface{}
	fakePayload := []model.EmployeeReturn{
		model.EmployeeReturn{ID: "1", FirstName: "Leon", LastName: "Ceasar", Email: "leon.ceasar@mail.com", Department: "fakeDepartment"},
		model.EmployeeReturn{ID: "1", FirstName: "Leon", LastName: "Ceasar", Email: "leon.ceasar@mail.com", Department: "fakeDepartment"},
		model.EmployeeReturn{ID: "1", FirstName: "Leon", LastName: "Ceasar", Email: "leon.ceasar@mail.com", Department: "fakeDepartment"},
	}
	fakeNilPayload := []model.EmployeeReturn{}
	fakeDecodeErr := errors.New("Decode went wrong")
	fakeNoResultErr := errors.New("No results could be found to your query")
	var tests = []struct {
		hasDecodeErr bool
		hasNoPayload bool
		payload      []model.EmployeeReturn
		err          error
	}{
		{false, false, fakePayload, nil},
		{true, false, fakeNilPayload, fakeDecodeErr},
		{false, true, fakeNilPayload, nil},
	}

	for _, tt := range tests {
		fakeDb.GetEmployeesByDepartmentReturns(tt.payload, tt.err)
		serviceInstance := service.NewEmployeeService(fakeDb)

		if !tt.hasNoPayload && !tt.hasDecodeErr && tt.err == nil {
			actualResult, actualErr := serviceInstance.GetEmployeesDepartmentFilter(tt.payload[0].Department)
			assert.Equal(t, fakePayload, actualResult)
			assert.Equal(t, tt.err, actualErr)
		}
		if tt.hasDecodeErr {
			actualResult, actualErr := serviceInstance.GetEmployeesDepartmentFilter("fakeDepartment")
			assert.Equal(t, tt.payload, actualResult)
			assert.Equal(t, tt.err, actualErr)
		}

		if tt.hasNoPayload {
			actualResult, actualErr := serviceInstance.GetEmployeesDepartmentFilter("fakeDepartment")
			assert.Equal(t, tt.payload, actualResult)
			assert.Equal(t, fakeNoResultErr, actualErr)
		}
	}
}
