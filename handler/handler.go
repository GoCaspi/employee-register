package handler

import (
	"crypto/subtle"
	"encoding/json"
	"example-project/cache"
	"example-project/model"
	"example-project/utility"
	"fmt"
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
	GetEmployeesDepartmentFilter(department string) ([]model.EmployeeReturn, error)
	AddShift(emp model.Employee, shift model.Shift) ([]model.Shift, error)
	GetRoster(employees []model.EmployeeReturn, week int) (map[string]map[string]model.Workload, error)
}

var MyCacheMap = cache.NewCacheMap{}

const noTokenErr = "No token is provided. Please login in and provide a token"
const noEmployeeFound = "Cannot find an employee to that id!"
const invalidPayloadMsg = "invalid payload"

const clientID = "69678bb4a1b8a0c2462f"
const clientSecret = "cad250266a5613152d5a9ea64e70429545855782"

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
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

		pageInt = 1
		limitInt = 1000000 * 100000

		response, _ := handler.ServiceInterface.GetPaginatedEmployees(pageInt, limitInt)

		c.JSON(http.StatusOK, response)
	}

}

func (handler Handler) OAuthRedirectHandler(context *gin.Context) {
	code := context.Query("code")
	reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	req.Header.Set("accept", "application/json")
	if err != nil {
		fmt.Println(err)
	}
	httpClient := http.Client{}
	// Send out the HTTP request
	res, _ := httpClient.Do(req)

	//	fmt.Println(res)
	var t OAuthAccessResponse
	if err = json.NewDecoder(res.Body).Decode(&t); err != nil {
		context.AbortWithStatusJSON(402, "Couldnt fetch access token")
	}

	githubUserId := uuid.New().String()

	cache.AddToCacheMap(githubUserId, t.AccessToken, MyCacheMap)

	guestMsg := "Success! Your Guest-Id is :" + githubUserId + " and your guest-token is: " + t.AccessToken

	context.JSON(200, guestMsg)
	//	githubData := getGithubData(t.AccessToken)

	//	fmt.Println(githubData)

	//	context.JSON(200, githubData)
}

func (handler Handler) OAuthStarterHandler(context *gin.Context) {
	context.JSON(200, "https://github.com/login/oauth/authorize?client_id=69678bb4a1b8a0c2462f")
}

/*
func getGithubData(accessToken string) string {
	// Get request to a set URL
	req, reqerr := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if reqerr != nil {
		log.Panic("API Request creation failed")
	}

	// Set the Authorization header before sending the request
	// Authorization: token XXXXXXXXXXXXXXXXXXXXXXXXXXX
	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	// Make the request
	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}

	// Read the response as a byte slice
	respbody, _ := ioutil.ReadAll(resp.Body)

	// Convert byte slice to string and return
	return string(respbody)
}

*/
func (handler Handler) DepartmentFilter(context *gin.Context) {
	department, depOk := context.GetQuery("department")
	if !depOk {
		noQueryError := "No department was given in the query parameter!"
		context.AbortWithStatusJSON(404, gin.H{
			"errorMessage": noQueryError,
		})
		return
	}

	response, err := handler.ServiceInterface.GetEmployeesDepartmentFilter(department)
	if err != nil {
		context.AbortWithStatusJSON(404, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	context.JSON(200, response)
	return

}

func (handler Handler) AddShift(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if !ok {
		noQueryError := "No Id was submitted. Please add an id to your query"
		context.AbortWithStatusJSON(404, gin.H{
			"errorMessage": noQueryError,
		})
		return
	}

	var shiftPayload model.Shift
	err := context.ShouldBindJSON(&shiftPayload)
	if err != nil {
		context.AbortWithStatusJSON(400, gin.H{
			"errorMessage": "Bad payload",
		})
		return
	}

	employee := handler.ServiceInterface.GetEmployeeById(id)

	response, err := handler.ServiceInterface.AddShift(employee, shiftPayload)
	if err != nil {
		context.AbortWithStatusJSON(404, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	context.JSON(200, response)

}

func (handler Handler) GetDutyRoster(context *gin.Context) {
	department, depOk := context.GetQuery("department")
	if !depOk {
		noDepartmentQuery := "No department was specified in the query."
		context.AbortWithStatusJSON(404, gin.H{
			"errorMessage": noDepartmentQuery,
		})
		return
	}
	week, weekOk := context.GetQuery("week")
	weekInt, strConvErr := strconv.Atoi(week)
	if !weekOk {
		noWeekQuery := "No week was specified in the query."
		context.AbortWithStatusJSON(404, gin.H{
			"errorMessage": noWeekQuery,
		})
		return
	}

	if strConvErr != nil {
		context.AbortWithStatusJSON(400, gin.H{
			"errorMessage": strConvErr.Error(),
		})
	}

	departmentEmployees, err := handler.ServiceInterface.GetEmployeesDepartmentFilter(department)

	if err != nil {
		context.AbortWithStatusJSON(404, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	roster, err := handler.ServiceInterface.GetRoster(departmentEmployees, weekInt)

	if err != nil {
		context.AbortWithStatusJSON(404, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	context.JSON(200, roster)
}

/*
[
{
"week": 1,
"duties": {
"Monday": {
"duty": "Cleaninig",
"start": "2006-01-02T15:04:05Z",
"end": "2006-01-02T17:04:05Z",
"total": 7200000000000
}
}
}
]

*/
