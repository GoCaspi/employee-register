package handler

import (
	"example-project/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ServiceInterface
type ServiceInterface interface {
	CreateEmployees(employees []model.Employee) (interface{}, error)
	GetEmployeeById(id string) model.Employee
}

type Handler struct {
	ServiceInterface ServiceInterface
}

func NewHandler(serviceInterface ServiceInterface) Handler {
	return Handler{
		ServiceInterface: serviceInterface,
	}
}

func (handler Handler) CreateEmployeeHandler(c *gin.Context) {
	var payLoad model.Payload
	err := c.ShouldBindJSON(&payLoad)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "invalid payload",
		})
		return
	}

	response, err := handler.ServiceInterface.CreateEmployees(payLoad.Employees)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Create Employees failed due to a database error",
		})
		return
	}

	c.JSON(200, response)
}

func (handler Handler) GetEmployeeHandler(c *gin.Context) {
	pathParam, ok := c.Params.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "id is not given",
		})
		return
	}

	response := handler.ServiceInterface.GetEmployeeById(pathParam)
	c.JSON(http.StatusOK, response)
}
