package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . HandlerInterface
type HandlerInterface interface {
	CreateEmployeeHandler(c *gin.Context)
	GetEmployeeHandler(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	ValidateToken(c *gin.Context)
	DeleteByIdHandler(c *gin.Context)
}

var Handler HandlerInterface

const clientID = "69678bb4a1b8a0c2462f"
const clientSecret = "cad250266a5613152d5a9ea64e70429545855782"

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

func CreateRoutes(group *gin.RouterGroup) {

	group.POST("/Login", Handler.Login)
	group.POST("/Logout", Handler.Logout)
	group.POST("/register", Handler.CreateEmployeeHandler)
	group.GET("/tester", func(context *gin.Context) {
		context.JSON(200, "https://github.com/login/oauth/authorize?client_id=69678bb4a1b8a0c2462f")
	})

	group.GET("/authRedirect", func(context *gin.Context) {
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

		fmt.Println(res)

		//	code := context.Query("code")
		context.JSON(200, code)
	})

	route := group.Group("/employee")
	route.Use(Handler.ValidateToken)
	route.GET("/:id/get", Handler.GetEmployeeHandler)
	route.POST("/create", Handler.CreateEmployeeHandler)
	route.DELETE("/:id/delete", Handler.DeleteByIdHandler)
}
