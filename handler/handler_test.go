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

func TestGetEmployeeHandler_Return_valid_status_code(t *testing.T) {
	responseRecoder := httptest.NewRecorder()

	fakeContest, _ := gin.CreateTestContext(responseRecoder)
	fakeContest.Params = append(fakeContest.Params, gin.Param{Key: "id", Value: "1"})

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.GetEmployeeByIdReturns(model.Employee{
		ID:        "1",
		FirstName: "Joe",
	})

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.GetEmployeeHandler(fakeContest)

	assert.Equal(t, http.StatusOK, responseRecoder.Code)

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
