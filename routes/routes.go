package routes

import (
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
	OauthUrlGetter(c *gin.Context)
	OauthCallback(c *gin.Context)
}

var Handler HandlerInterface

func CreateRoutes(group *gin.RouterGroup) {

	group.POST("/Login", Handler.Login)
	group.POST("/Logout", Handler.Logout)
	group.POST("/register", Handler.CreateEmployeeHandler)
	group.GET("/registeralt", Handler.OauthUrlGetter)
	group.GET("/oauthCallback", Handler.OauthCallback)
	group.StaticFS("/static", http.Dir("./dist"))
	route := group.Group("/employee")
	route.Use(Handler.ValidateToken)
	route.GET("/:id/get", Handler.GetEmployeeHandler)
	route.POST("/create", Handler.CreateEmployeeHandler)
	route.DELETE("/:id/delete", Handler.DeleteByIdHandler)
}
