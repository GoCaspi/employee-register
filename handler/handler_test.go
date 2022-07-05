package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"example-project/handler"
	"example-project/handler/handlerfakes"
	"example-project/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetEmployeeById(t *testing.T) {
	fakeErr := errors.New("")
	mockEmp := model.Employee{
		ID:        "1",
		FirstName: "Joe",
	}
	var tests = []struct {
		badParams  bool
		serviceErr bool
		employee   model.Employee
	}{
		{false, false, mockEmp},
		{true, false, mockEmp},
		{false, true, mockEmp},
	}

	for _, tt := range tests {

		if !tt.serviceErr && !tt.badParams {
			responseRecoder := httptest.NewRecorder()

			fakeContest, _ := gin.CreateTestContext(responseRecoder)
			fakeContest.Params = append(fakeContest.Params, gin.Param{Key: "id", Value: tt.employee.ID})
			fakeService := &handlerfakes.FakeServiceInterface{}
			fakeService.GetEmployeeByIdReturns(tt.employee, nil)
			handlerInstance := handler.NewHandler(fakeService)
			handlerInstance.GetEmployeeHandler(fakeContest)
			assert.Equal(t, http.StatusOK, responseRecoder.Code)
		}

		if tt.badParams {
			responseRecoder := httptest.NewRecorder()

			fakeContest, _ := gin.CreateTestContext(responseRecoder)
			fakeContest.Params = append(fakeContest.Params, gin.Param{Key: "i", Value: "1"})
			fakeService := &handlerfakes.FakeServiceInterface{}
			fakeService.GetEmployeeByIdReturns(tt.employee, nil)
			handlerInstance := handler.NewHandler(fakeService)
			handlerInstance.GetEmployeeHandler(fakeContest)
			assert.Equal(t, http.StatusBadRequest, responseRecoder.Code)
		}

		if tt.serviceErr {
			responseRecoder := httptest.NewRecorder()

			fakeContest, _ := gin.CreateTestContext(responseRecoder)
			fakeContest.Params = append(fakeContest.Params, gin.Param{Key: "id", Value: tt.employee.ID})
			fakeService := &handlerfakes.FakeServiceInterface{}
			fakeService.GetEmployeeByIdReturns(tt.employee, fakeErr)
			handlerInstance := handler.NewHandler(fakeService)
			handlerInstance.GetEmployeeHandler(fakeContest)
			assert.Equal(t, http.StatusBadRequest, responseRecoder.Code)
		}

	}

}

func TestHandler_CreateEmployeeHandler(t *testing.T) {

	jsonPayload := `{
  "employees": [
    {
      "id": "1",
      "first_name": "John",
      "last_name": "Kenn",
      "email": "john@gmail.com"
    },
    {
      "id": "2",
      "first_name": "Maria",
      "last_name": "gonjaless",
      "email": "maria@gmail.com"
    },
    {
      "id": "3",
      "first_name": "Lora",
      "last_name": "kai",
      "email": "lora@gmail.com"
    }
  ]
}`
	fakeErr := errors.New("Db error triggerd")
	fakeInterface := struct {
		InsertedIDs int
	}{
		1,
	}

	var tests = []struct {
		json       string
		hasDbError bool
		fakeError  error
		badPayload bool
	}{
		{jsonPayload, false, nil, false},
		{jsonPayload, false, nil, true},
		{jsonPayload, true, fakeErr, false},
	}

	for _, tt := range tests {
		if !tt.hasDbError && !tt.badPayload {
			responseRecoder := httptest.NewRecorder()
			fakeContest, _ := gin.CreateTestContext(responseRecoder)
			fakeService := &handlerfakes.FakeServiceInterface{}

			var mockPayload model.Payload
			json.Unmarshal([]byte(tt.json), &mockPayload)
			body := bytes.NewBufferString(tt.json)

			fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/employees", body)
			fakeService.CreateEmployeesReturns(fakeInterface, nil)
			handlerInstance := handler.NewHandler(fakeService)
			handlerInstance.CreateEmployeeHandler(fakeContest)

			assert.Equal(t, http.StatusOK, responseRecoder.Code)
		}

		if tt.badPayload {
			responseRecoder := httptest.NewRecorder()
			fakeContest, _ := gin.CreateTestContext(responseRecoder)
			fakeService := &handlerfakes.FakeServiceInterface{}

			fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/employees", nil)
			fakeService.CreateEmployeesReturns(fakeInterface, nil)
			handlerInstance := handler.NewHandler(fakeService)
			handlerInstance.CreateEmployeeHandler(fakeContest)

			assert.Equal(t, http.StatusBadRequest, responseRecoder.Code)
		}

		if tt.hasDbError {
			responseRecoder := httptest.NewRecorder()
			fakeContest, _ := gin.CreateTestContext(responseRecoder)
			fakeService := &handlerfakes.FakeServiceInterface{}

			var mockPayload model.Payload
			json.Unmarshal([]byte(tt.json), &mockPayload)
			body := bytes.NewBufferString(tt.json)

			fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/employees", body)
			fakeService.CreateEmployeesReturns(fakeInterface, tt.fakeError)
			handlerInstance := handler.NewHandler(fakeService)
			handlerInstance.CreateEmployeeHandler(fakeContest)

			assert.Equal(t, http.StatusBadRequest, responseRecoder.Code)
		}
	}

}
