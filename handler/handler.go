package handler

import (
	"crypto/subtle"
	"example-project/cache"
	"example-project/model"
	"example-project/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"

	//	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"github.com/google/uuid"
	"net/http"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ServiceInterface
type ServiceInterface interface {
	CreateEmployees(employees []model.Employee) (interface{}, error)
	GetEmployeeById(id string) model.Employee
	DeleteEmployee(id string) (interface{}, error)
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

func (handler Handler) OauthUrlGetter(c *gin.Context) {
	service, serviceOK := c.GetQuery("service")
	if serviceOK == false {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errorMessage": "service query wasn't supplied",
		})
		return
	}
	if service == "github" {
		state := time.Now().Unix()
		scopeList := "read:user%20user:email"
		githubURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%v&scope=%v$state=%d", os.Getenv("OAUTH_GITHUB_CLIENT_ID"), scopeList, state)
		c.JSON(http.StatusOK, gin.H{
			"service": "github",
			"url":     githubURL,
		})
	}
}

func (handler Handler) OauthCallback(c *gin.Context) {
	code, codeOK := c.GetQuery("code")
	fmt.Println(code)
	if codeOK == true {
		c.Redirect(http.StatusMovedPermanently, "/static/success.html")
		var id = uuid.New()
		cache.AddToCacheMap(id.String(), code, MyCacheMap)
		emp := []model.Employee{
			model.Employee{ID: id.String(), FirstName: "Thomas", LastName: "Müller", Email: "thomas.müller@gmail.com", Auth: model.HashedAuth{}},
		}
		result, err := handler.ServiceInterface.CreateEmployees(emp)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errorMessage": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"guestToken": code,
			"guestID":    id.String(),
			"results":    result,
		})
		return
	} else {
		c.Redirect(http.StatusMovedPermanently, "/static/fail.html")
	}
}
