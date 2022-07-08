package handler_test

import (
	"example-project/cache"
	"example-project/handler"
	"example-project/handler/handlerfakes"
	"example-project/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func TestHandler_Logout(t *testing.T) {
	uuid := uuid.New()
	uuidString := uuid.String()

	var tests = []struct {
		isLoggedIn     bool
		id             string
		keyIsMissing   bool
		expectedStatus int
	}{
		{true, "1", false, 400},
		{true, "1", true, 400},
		{false, "1", false, 200},
		{false, "1", true, 400},
	}

	for _, tt := range tests {
		responseRecoder := httptest.NewRecorder()
		fakeContest, _ := gin.CreateTestContext(responseRecoder)
		fakeService := &handlerfakes.FakeServiceInterface{}
		handlerInstance := handler.NewHandler(fakeService)
		if tt.isLoggedIn && !tt.keyIsMissing {
			handler.MyCacheMap = cache.AddToCacheMap(tt.id, uuidString, handler.MyCacheMap)
			fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/Logout?id="+tt.id, nil)
			handlerInstance.Logout(fakeContest)
			assert.Equal(t, 200, responseRecoder.Code)

		}
		if tt.keyIsMissing && tt.isLoggedIn {
			handler.MyCacheMap = cache.AddToCacheMap(tt.id, uuidString, handler.MyCacheMap)
			fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/Logout?+tt.id", nil)
			handlerInstance.Logout(fakeContest)
			assert.Equal(t, 400, responseRecoder.Code)

		}

		if !tt.isLoggedIn {
			handler.MyCacheMap = cache.RemoveFromCacheMap(tt.id, handler.MyCacheMap)
			fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/Logout?id="+tt.id, nil)
			handlerInstance.Logout(fakeContest)
			assert.Equal(t, 400, responseRecoder.Code)

		}
	}

}

func TestHandler_ValidateToken(t *testing.T) {
	uuid := uuid.New()
	uuidString := uuid.String()

	var tests = []struct {
		tokenIsValid   bool
		token          string
		fakeId         string
		tokenIsPresent bool
	}{
		{true, uuidString, "1", true},
		{false, uuidString, "1", true},
		{false, uuidString, "1", false},
	}

	for _, tt := range tests {
		fakeService := &handlerfakes.FakeServiceInterface{}
		handlerInstance := handler.NewHandler(fakeService)

		responseRecoder := httptest.NewRecorder()
		fakeContest, _ := gin.CreateTestContext(responseRecoder)
		if tt.tokenIsValid && tt.tokenIsPresent {
			handler.MyCacheMap = cache.AddToCacheMap(tt.fakeId, uuidString, handler.MyCacheMap)
			fakeContest.Request = httptest.NewRequest("GET", "http://localhost:9090/token", nil)
			fakeContest.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tt.token))

			handlerInstance.ValidateToken(fakeContest)
			assert.Equal(t, 200, responseRecoder.Code)
		}
		if !tt.tokenIsValid && tt.tokenIsPresent {
			handler.MyCacheMap = cache.RemoveFromCacheMap(tt.fakeId, handler.MyCacheMap)
			fakeContest.Request = httptest.NewRequest("GET", "http://localhost:9090/token", nil)
			fakeContest.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tt.token))

			handlerInstance.ValidateToken(fakeContest)
			assert.Equal(t, 401, responseRecoder.Code)
		}

		if !tt.tokenIsPresent {
			handler.MyCacheMap = cache.RemoveFromCacheMap(tt.fakeId, handler.MyCacheMap)
			fakeContest.Request = httptest.NewRequest("GET", "http://localhost:9090/token", nil)

			handlerInstance.ValidateToken(fakeContest)
			assert.Equal(t, 403, responseRecoder.Code)
		}

	}
}

func TestHandler_Login(t *testing.T) {

}

func TestHandler_DoUserExist(t *testing.T) {
	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeEmployees := []model.Employee{
		model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
		model.Employee{ID: "200", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
	}

	fakeEmployeesDuplication := []model.Employee{
		model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
		model.Employee{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
	}

	var tests = []struct {
		Result   []model.Employee
		expected bool
		isEmpty  bool
	}{
		{fakeEmployeesDuplication, true, false},
		{fakeEmployees, false, false},
		{[]model.Employee{fakeEmployees[0]}, false, false},
		{fakeEmployees, false, true},
	}
	for _, tt := range tests {

		if tt.isEmpty {
			var emptyEmployee model.Employee
			fakeService.GetEmployeeByIdReturns(emptyEmployee)
			handlerInstance := handler.NewHandler(fakeService)
			boolResult, _ := handlerInstance.DoUserExist(tt.Result)
			assert.Equal(t, tt.expected, boolResult)
		} else {
			if tt.expected {
				handlerInstance := handler.NewHandler(fakeService)
				boolResult, _ := handlerInstance.DoUserExist(tt.Result)
				assert.Equal(t, tt.expected, boolResult)
			} else {
				handlerInstance := handler.NewHandler(fakeService)
				boolResult, _ := handlerInstance.DoUserExist(tt.Result)
				assert.Equal(t, tt.expected, boolResult)
			}
		}
	}
}
