package handler

import (
	"crypto/subtle"
	"example-project/cache"
	"example-project/model"
	"example-project/utility"
	"github.com/gin-gonic/gin"
	"strconv"

	//	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"github.com/google/uuid"
	"net/http"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ServiceInterface
type ServiceInterface interface {
	CreateEmployees(employees []model.Employee) (interface{}, error)
	GetEmployeeById(id string) model.Employee
	DeleteEmployee(id string) (interface{}, error)
	GetPaginatedEmployees(page int, limit int) (model.PaginatedPayload, error)
}

var MyCacheMap = cache.NewCacheMap{}

const noTokenErr = "No token is provided. Please login in and provide a token"
const noEmployeeFound = "Cannot find an employee to that id!"
const invalidPayloadMsg = "invalid payload"

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

	employees := payLoad.Employees
	hashedEmployees := utility.HashEmployees(employees)
	index, _ := handler.DoUserExist(hashedEmployees)
	if index {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "There cannot be duplicated Ids!",
		})
		return
	}
	response, err := handler.ServiceInterface.CreateEmployees(hashedEmployees)

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
	employee := model.EmployeeReturn{
		ID:        response.ID,
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Email:     response.Email,
	}
	c.JSON(http.StatusOK, employee)
}

func (handler Handler) Login(c *gin.Context) {
	id, keyIsPresent := c.GetQuery("id")
	errMsg := noEmployeeFound
	if !keyIsPresent {
		c.AbortWithStatusJSON(400, errMsg)
		return
	}
	employee := handler.ServiceInterface.GetEmployeeById(id)
	//	if err != nil {
	//		c.AbortWithStatusJSON(400, noEmployeeFound)
	//		return
	//	}
	var payLoad model.Auth
	err := c.ShouldBindJSON(&payLoad)
	errMsg = invalidPayloadMsg
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg)
		return
	}
	if err == nil {

		actualUser, actualPwd := utility.HashUsernameAndPassword(payLoad).Username, utility.HashUsernameAndPassword(payLoad).Password
		validUser, validPwd := employee.Auth.Username, employee.Auth.Password

		usernameMatch := subtle.ConstantTimeCompare(actualUser[:], validUser[:]) == 1
		passwordMatch := subtle.ConstantTimeCompare(actualPwd[:], validPwd[:]) == 1

		if usernameMatch && passwordMatch {
			uuid := uuid.New()
			uuidString := uuid.String()
			successMsg := "Success! Your Token is: " + uuidString
			MyCacheMap = cache.AddToCacheMap(employee.ID, uuidString, MyCacheMap)

			c.JSON(http.StatusOK, successMsg)
			return
		}
	}
	errMsg = "The username or password is wrong"
	c.AbortWithStatusJSON(401, errMsg)
	c.Writer.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	return
}

func (handler Handler) DoUserExist(emp []model.Employee) (bool, []model.Employee) {
	var idList []string
	var errorEmployees []model.Employee

	for _, employee := range emp {
		response := handler.ServiceInterface.GetEmployeeById(employee.ID)
		if len(response.ID) != 0 {
			errorEmployees = append(errorEmployees, employee)
		} else {
			idList = append(idList, employee.ID)
			var idCount int = 0
			for _, id := range idList {
				if id == employee.ID {
					idCount++
				}
			}
			if idCount >= 2 {
				errorEmployees = append(errorEmployees, employee)
			}
		}
	}
	if len(errorEmployees) != 0 {
		return true, errorEmployees
	} else {
		return false, nil
	}
}

func (handler Handler) ValidateToken(c *gin.Context) {

	if len(c.Request.Header.Values("Authorization")) < 1 {
		c.AbortWithStatusJSON(403, noTokenErr)
		return
	}
	reqToken := utility.GetBearerToken(c)

	tokenIsValid := cache.TokenIsInMap(reqToken, MyCacheMap)
	if !tokenIsValid {
		c.AbortWithStatusJSON(401, noTokenErr)

	} else {
		return
	}
}

func (handler Handler) Logout(c *gin.Context) {
	id, keyIsPresent := c.GetQuery("id")
	errMsg := noEmployeeFound
	if !keyIsPresent {
		c.AbortWithStatusJSON(400, errMsg)
		return
	}

	tokenIsInCache := cache.IdIsInMap(id, MyCacheMap)
	successMessage := "Logut successfull. Your token is no longer valid."
	failMessage := "The provided token is not valid. Please login to generate a valid token."

	if tokenIsInCache {
		MyCacheMap = cache.RemoveFromCacheMap(id, MyCacheMap)
		c.JSON(200, successMessage)
	} else {
		c.JSON(400, failMessage)
	}
}

func (handler Handler) DeleteByIdHandler(c *gin.Context) {
	pathParam, ok := c.Params.Get("id")

	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{

			"errorMessage": "id is not found",
		})
		return
	}
	response, err := handler.ServiceInterface.DeleteEmployee(pathParam)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (handler Handler) GetAllEmployeesHandler(c *gin.Context) {
	pages, pageOk := c.GetQuery("page")
	limit, limitOk := c.GetQuery("limit")
	pageInt, pageErr := strconv.Atoi(pages)
	limitInt, limitErr := strconv.Atoi(limit)
	if pageOk && limitOk {
		if pageOk && limitOk && pageErr == nil && limitErr == nil {

			response, err := handler.ServiceInterface.GetPaginatedEmployees(pageInt, limitInt)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"errorMessage": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, response)
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errorMessage": "queries are invalid, please check or remove them",
			})
			return
		}
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "This is a paginated endpoint, please submit a limit and a page",
		})
		return
	}

}
