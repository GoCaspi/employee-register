package handler_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"example-project/cache"
	"example-project/handler"
	"example-project/handler/handlerfakes"
	"example-project/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
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

func TestGetEmployeeHandler_Return_Invalid_status_code_woringQueryParams(t *testing.T) {
	responseRecoder := httptest.NewRecorder()

	fakeContest, _ := gin.CreateTestContext(responseRecoder)
	fakeContest.Params = append(fakeContest.Params, gin.Param{Key: "ld", Value: "1"})

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.GetEmployeeByIdReturns(model.Employee{
		ID:        "1",
		FirstName: "Joe",
	})

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.GetEmployeeHandler(fakeContest)

	assert.Equal(t, http.StatusBadRequest, responseRecoder.Code)

}
func TestHandler_DeleteByIdHandler(t *testing.T) {
	ResponseRecorder := httptest.NewRecorder()

	fakeContext, _ := gin.CreateTestContext(ResponseRecorder)
	fakeContext.Params = append(fakeContext.Params, gin.Param{Key: "id", Value: "1"})
	fakeContext.Params = append(fakeContext.Params, gin.Param{Key: "MIl", Value: "x"})

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.DeleteEmployeeReturns(&mongo.DeleteResult{DeletedCount: 1}, nil)

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.DeleteByIdHandler(fakeContext)

	assert.Equal(t, 200, ResponseRecorder.Code)

}
func TestHandler_Runpassparameter(t *testing.T) {
	ResponseRecorder := httptest.NewRecorder()

	fakeContext, _ := gin.CreateTestContext(ResponseRecorder)
	fakeContext.Params = append(fakeContext.Params, gin.Param{Key: "MIl", Value: "x"})

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.DeleteEmployeeReturns(&mongo.DeleteResult{DeletedCount: 1}, nil)

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.DeleteByIdHandler(fakeContext)

	assert.Equal(t, 400, ResponseRecorder.Code)
}
func TestHandler_EmployeeDeleted(t *testing.T) {
	ResponseRecorder := httptest.NewRecorder()

	fakeContext, _ := gin.CreateTestContext(ResponseRecorder)
	fakeContext.Params = append(fakeContext.Params, gin.Param{Key: "id", Value: "1"})

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakeService.DeleteEmployeeReturns(&mongo.DeleteResult{DeletedCount: 1}, errors.New("MIL"))

	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.DeleteByIdHandler(fakeContext)

	assert.Equal(t, 400, ResponseRecorder.Code)
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
func TestHandler_Login_Return_valid_status_code(t *testing.T) {
	responseRecoder := httptest.NewRecorder()

	jsonPayload := `{
		"username": "fakeUser",
		"password": "fakePwd"
		    }`

	var mockAuth model.Auth
	json.Unmarshal([]byte(jsonPayload), &mockAuth)
	body := bytes.NewBufferString(jsonPayload)

	fakeContest, _ := gin.CreateTestContext(responseRecoder)
	fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/auth/Login?id=1", body)

	fakeAuth := model.Auth{Password: "fakePwd", Username: "fakeUser"}

	usernameHash := sha256.Sum256([]byte(fakeAuth.Username))
	passwordHash := sha256.Sum256([]byte(fakeAuth.Password))

	fakeService := &handlerfakes.FakeServiceInterface{}
	auth := model.HashedAuth{Username: usernameHash, Password: passwordHash}
	fakeService.GetEmployeeByIdReturns(model.Employee{
		ID:        "1",
		FirstName: "Joe",
		LastName:  "Doe",
		Email:     "john@doe.de",
		Auth:      auth,
	})
	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.Login(fakeContest)
	assert.Equal(t, http.StatusOK, responseRecoder.Code)
}

func TestHandler_Login_Return_InvalidStatusCode_EmployeeToGivenIDNotFoundInDatabase(t *testing.T) {
	responseRecoder := httptest.NewRecorder()

	jsonPayload := `{
		"username": "fakeUser",
		"password": "fakePwd"
		    }`

	var mockAuth model.Auth
	json.Unmarshal([]byte(jsonPayload), &mockAuth)
	body := bytes.NewBufferString(jsonPayload)

	fakeContest, _ := gin.CreateTestContext(responseRecoder)
	fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/auth/Login?id=1", body)

	fakeAuth := model.Auth{Password: "fakePwd", Username: "fakeUser"}

	usernameHash := sha256.Sum256([]byte(fakeAuth.Username))
	passwordHash := sha256.Sum256([]byte(fakeAuth.Password))

	fakeService := &handlerfakes.FakeServiceInterface{}
	auth := model.HashedAuth{Username: usernameHash, Password: passwordHash}
	fakeService.GetEmployeeByIdReturns(model.Employee{
		ID:        "1",
		FirstName: "Joe",
		LastName:  "Doe",
		Email:     "john@doe.de",
		Auth:      auth,
	})

	expectedErrorMsg := ""
	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.Login(fakeContest)

	assert.Contains(t, responseRecoder.Body.String(), expectedErrorMsg)
}

func TestHandler_Login_Return_InvalidStatusCode_InvalidPayloadInPostRequest(t *testing.T) {
	responseRecoder := httptest.NewRecorder()

	jsonPayload := `
		"fakename": "fakeUser",
		"password": "fakePwd"
		    }`

	var mockAuth model.Auth
	json.Unmarshal([]byte(jsonPayload), &mockAuth)
	body := bytes.NewBufferString(jsonPayload)

	fakeContest, _ := gin.CreateTestContext(responseRecoder)
	fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/auth/Login?id=1", body)

	fakeAuth := model.Auth{Password: "fakePwd", Username: "fakeUser"}

	usernameHash := sha256.Sum256([]byte(fakeAuth.Username))
	passwordHash := sha256.Sum256([]byte(fakeAuth.Password))

	fakeService := &handlerfakes.FakeServiceInterface{}
	auth := model.HashedAuth{Username: usernameHash, Password: passwordHash}
	//	fakeError := errors.New("")
	fakeService.GetEmployeeByIdReturns(model.Employee{
		ID:        "1",
		FirstName: "Joe",
		LastName:  "Doe",
		Email:     "john@doe.de",
		Auth:      auth,
	})

	expectedErrorMsg := ""
	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.Login(fakeContest)

	assert.Contains(t, responseRecoder.Body.String(), expectedErrorMsg)
}
func TestHandler_Login_Return_InvalidStatusCode_PostedUsernameAndPassword_DontMatchSavedDataInDatabase(t *testing.T) {
	responseRecoder := httptest.NewRecorder()

	jsonPayload := `{
		"username": "fakeUser",
		"password": "Pwd"
		    }`

	var mockAuth model.Auth
	json.Unmarshal([]byte(jsonPayload), &mockAuth)
	body := bytes.NewBufferString(jsonPayload)

	fakeContest, _ := gin.CreateTestContext(responseRecoder)
	fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/auth/Login?id=1", body)

	fakeAuth := model.Auth{Password: "fakePwd", Username: "fakeUser"}

	usernameHash := sha256.Sum256([]byte(fakeAuth.Username))
	passwordHash := sha256.Sum256([]byte(fakeAuth.Password))

	fakeService := &handlerfakes.FakeServiceInterface{}
	auth := model.HashedAuth{Username: usernameHash, Password: passwordHash}
	fakeService.GetEmployeeByIdReturns(model.Employee{
		ID:        "1",
		FirstName: "Joe",
		LastName:  "Doe",
		Email:     "john@doe.de",
		Auth:      auth,
	})

	expectedErrorMsg := "The username or password is wrong"
	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.Login(fakeContest)

	assert.Contains(t, responseRecoder.Body.String(), expectedErrorMsg)
}

func TestHandler_Login_Return_InvalidStatusCode_QueryKeyIsWrong(t *testing.T) {
	responseRecoder := httptest.NewRecorder()

	jsonPayload := `{
		"username": "fakeUser",
		"password": "Pwd"
		    }`

	var mockAuth model.Auth
	json.Unmarshal([]byte(jsonPayload), &mockAuth)
	body := bytes.NewBufferString(jsonPayload)

	fakeContest, _ := gin.CreateTestContext(responseRecoder)
	fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/auth/Login?1", body)

	fakeAuth := model.Auth{Password: "fakePwd", Username: "fakeUser"}

	usernameHash := sha256.Sum256([]byte(fakeAuth.Username))
	passwordHash := sha256.Sum256([]byte(fakeAuth.Password))

	fakeService := &handlerfakes.FakeServiceInterface{}
	auth := model.HashedAuth{Username: usernameHash, Password: passwordHash}
	fakeService.GetEmployeeByIdReturns(model.Employee{
		ID:        "1",
		FirstName: "Joe",
		LastName:  "Doe",
		Email:     "john@doe.de",
		Auth:      auth,
	})

	expectedErrorMsg := ""
	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.Login(fakeContest)

	assert.Contains(t, responseRecoder.Body.String(), expectedErrorMsg)
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

func TestHandler_CreateEmployeeHandler(t *testing.T) {

	jsonPayload := `{
  "employees": [
    {
      "id": "1001",
      "first_name": "Jona",
      "last_name": "Miller",
      "email": "jona.millermail.com",
       "auth":  {
 "username":"Jona",
 "password":"pa55word"
}
    }
  ]
}`

	jsonPayloadDuplication := `{
  "employees": [
    {
      "id": "1001",
      "first_name": "Jona",
      "last_name": "Miller",
      "email": "jona.millermail.com",
       "auth":  {
 "username":"Jona",
 "password":"pa55word"
}
    },
 {
      "id": "1001",
      "first_name": "Jona",
      "last_name": "Miller",
      "email": "jona.millermail.com",
       "auth":  {
 "username":"Jona",
 "password":"pa55word"
}
    }
  ]
}`

	badJsonPayload := `
  oyees": [
    {
      "": "1001",
      "first_name": "Jona",
      "last_name": "Miller",
      "email": "jona.millermail.com",
       "auth":  {
 "username":"Jona",
 "password":"pa55word"
}
    }
  ]
}`
	/*
		var mockEmployee model.Employee
		json.Unmarshal([]byte(jsonPayload), &mockEmployee)
		body := bytes.NewBufferString(jsonPayload)

		fakeContest, _ := gin.CreateTestContext(responseRecoder)
		fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/register", body)
		//	expectedErrorMsg := "The username or password is wrong"
		fakeService := &handlerfakes.FakeServiceInterface{}
		handlerInstance := handler.NewHandler(fakeService)
		handlerInstance.CreateEmployeeHandler(fakeContest)
		assert.Equal(t, http.StatusOK, responseRecoder.Code)

	*/

	var tests = []struct {
		Payload        string
		badPayload     bool
		hasDuplication bool
	}{
		{jsonPayload, false, false},
		{badJsonPayload, true, false},
		{jsonPayloadDuplication, false, true},
	}
	for _, tt := range tests {
		responseRecoder := httptest.NewRecorder()
		var mockEmployee model.Employee
		json.Unmarshal([]byte(tt.Payload), &mockEmployee)
		body := bytes.NewBufferString(tt.Payload)
		fakeContest, _ := gin.CreateTestContext(responseRecoder)
		fakeService := &handlerfakes.FakeServiceInterface{}
		if !tt.badPayload && !tt.hasDuplication {
			fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/register", body)
			handlerInstance := handler.NewHandler(fakeService)
			handlerInstance.CreateEmployeeHandler(fakeContest)
			assert.Equal(t, http.StatusOK, responseRecoder.Code)
		}

		if tt.badPayload {
			fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/register", body)
			handlerInstance := handler.NewHandler(fakeService)
			handlerInstance.CreateEmployeeHandler(fakeContest)
			assert.Equal(t, http.StatusBadRequest, responseRecoder.Code)
		}

		if tt.hasDuplication {
			fakeContest.Request = httptest.NewRequest("POST", "http://localhost:9090/register", body)

			handlerInstance := handler.NewHandler(fakeService)

			handlerInstance.CreateEmployeeHandler(fakeContest)
			assert.Equal(t, http.StatusBadRequest, responseRecoder.Code)
		}
	}
}

func TestGetPaginatedEmployeesHandler_valid_request(t *testing.T) {
	fakeRecorder := httptest.NewRecorder()
	fakeContext, _ := gin.CreateTestContext(fakeRecorder)
	fakeContext.Request = httptest.NewRequest("POST", "http://localhost:9090/employee/get?page=1&limit=2", nil)

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakePaginatedPayload := model.PaginatedPayload{
		PageLimit: 2,
		Employees: []model.EmployeeReturn{
			model.EmployeeReturn{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
			model.EmployeeReturn{ID: "200", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
		},
	}
	fakeService.GetPaginatedEmployeesReturns(fakePaginatedPayload, nil)
	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.GetAllEmployeesHandler(fakeContext)
	assert.Equal(t, http.StatusOK, fakeRecorder.Code)

}

func TestGetPaginatedEmployeesHandler_invalid_request_pageiszero(t *testing.T) {
	fakeRecorder := httptest.NewRecorder()
	fakeContext, _ := gin.CreateTestContext(fakeRecorder)
	fakeContext.Request = httptest.NewRequest("POST", "http://localhost:9090/employee/get?page=0&limit=2", nil)

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakePaginatedPayload := model.PaginatedPayload{
		PageLimit: 0,
		Employees: []model.EmployeeReturn{
			{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
			{ID: "200", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
		},
	}
	invalidPageNumber := errors.New("invalid page number, page number can't be zero")
	fakeService.GetPaginatedEmployeesReturns(fakePaginatedPayload, invalidPageNumber)
	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.GetAllEmployeesHandler(fakeContext)
	assert.Equal(t, 400, fakeRecorder.Code)

}

func TestGetPaginatedEmployeesHandler_invalid_request_wrongquery(t *testing.T) {
	fakeRecorder := httptest.NewRecorder()
	fakeContext, _ := gin.CreateTestContext(fakeRecorder)
	fakeContext.Request = httptest.NewRequest("POST", "http://localhost:9090/employee/get?page=frgr&limit=2", nil)

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakePaginatedPayload := model.PaginatedPayload{
		PageLimit: 0,
		Employees: []model.EmployeeReturn{
			{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
			{ID: "200", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
		},
	}
	invalidPageNumber := errors.New("invalid page number, page number can't be zero")
	fakeService.GetPaginatedEmployeesReturns(fakePaginatedPayload, invalidPageNumber)
	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.GetAllEmployeesHandler(fakeContext)
	assert.Equal(t, 400, fakeRecorder.Code)

}

func TestGetPaginatedEmployeesHandler_Invalid_request_noQueryParamsGiven(t *testing.T) {
	fakeRecorder := httptest.NewRecorder()
	fakeContext, _ := gin.CreateTestContext(fakeRecorder)
	fakeContext.Request = httptest.NewRequest("POST", "http://localhost:9090/employee/get", nil)

	fakeService := &handlerfakes.FakeServiceInterface{}
	fakePaginatedPayload := model.PaginatedPayload{
		PageLimit: 2,
		Employees: []model.EmployeeReturn{
			model.EmployeeReturn{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
			model.EmployeeReturn{ID: "200", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com"},
		},
	}
	fakeService.GetPaginatedEmployeesReturns(fakePaginatedPayload, nil)
	handlerInstance := handler.NewHandler(fakeService)
	handlerInstance.GetAllEmployeesHandler(fakeContext)
	assert.Equal(t, 200, fakeRecorder.Code)

}

func TestHandler_DepartmentFilter(t *testing.T) {

	filterReturn := []model.EmployeeReturn{
		model.EmployeeReturn{ID: "100", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com", Department: "fakeDepartment"},
		model.EmployeeReturn{ID: "200", FirstName: "Test", LastName: "Tester", Email: "tester@gmail.com", Department: "fakeDepartment"},
	}

	filterEmptyReturn := []model.EmployeeReturn{}
	fakeError := errors.New("fake error triggered")
	var tests = []struct {
		noQueryParams bool
		serviceErr    bool
		expectedCode  int
		Return        []model.EmployeeReturn
		err           error
	}{
		{true, false, 404, filterEmptyReturn, fakeError},
		{false, true, 404, filterReturn, fakeError},
		{false, false, 200, filterReturn, nil},
	}

	for _, tt := range tests {
		fakeRecorder := httptest.NewRecorder()
		fakeContext, _ := gin.CreateTestContext(fakeRecorder)
		fakeContext.Request = httptest.NewRequest("GET", "http://localhost:9090/filter?department=fakeDepartment", nil)

		fakeService := &handlerfakes.FakeServiceInterface{}
		fakeService.GetEmployeesDepartmentFilterReturns(tt.Return, tt.err)

		if tt.noQueryParams {
			fakeContext.Request = httptest.NewRequest("GET", "http://localhost:9090/filter", nil)
			handlerInstance := handler.NewHandler(fakeService)
			handlerInstance.DepartmentFilter(fakeContext)
			assert.Equal(t, tt.expectedCode, fakeRecorder.Code)
		}

		if tt.serviceErr {
			handlerInstance := handler.NewHandler(fakeService)
			handlerInstance.DepartmentFilter(fakeContext)
			assert.Equal(t, tt.expectedCode, fakeRecorder.Code)

		}

		handlerInstance := handler.NewHandler(fakeService)
		handlerInstance.DepartmentFilter(fakeContext)
		assert.Equal(t, tt.expectedCode, fakeRecorder.Code)
	}
}

/*
func TestHandler_GetDutyRoster(t *testing.T) {
	departmentEmployeeReturn := []model.EmployeeReturn{
		model.EmployeeReturn{ID: "1", FirstName: "Elton", LastName: "Duck", Email: "elton.duck@mail.com",Department: "fakeDepartment"},
	}

	rosterReturn := map[string]map[string]model.Workload{}

	var tests = []struct{

	}
}

*/
