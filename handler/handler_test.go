package handler_test

import (
	"errors"
	"example-project/handler"
	"example-project/handler/handlerfakes"
	"example-project/model"
	"github.com/gin-gonic/gin"
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
