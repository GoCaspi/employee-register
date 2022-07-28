package routes

import (
	"github.com/gin-gonic/gin"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . HandlerInterface
type HandlerInterface interface {
	CreateEmployeeHandler(c *gin.Context)
	GetEmployeeHandler(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	ValidateToken(c *gin.Context)
	DeleteByIdHandler(c *gin.Context)
	GetAllEmployeesHandler(c *gin.Context)
	OAuthRedirectHandler(context *gin.Context)
	OAuthStarterHandler(context *gin.Context)
	UpdateById(context *gin.Context)
	ImplementrightManagementMaster(context *gin.Context)
	ImplementrightManagementBoss(context *gin.Context)
	ImplementrightManagementemployee(context *gin.Context)
	/*ImplementrightManagement(context *gin.Context)*/
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
	group.GET("/github", Handler.OAuthStarterHandler)
	group.PUT("/:id/update", Handler.UpdateById)
	group.GET("/:id/getTest", Handler.GetEmployeeHandler)

	group.GET("/authRedirect", Handler.OAuthRedirectHandler)

	route := group.Group("/employee")
	route.Use(Handler.ValidateToken)
	route.GET("/:id/get", Handler.ImplementrightManagementBoss, Handler.ImplementrightManagementemployee, Handler.GetEmployeeHandler)
	route.POST("/create", Handler.ImplementrightManagementBoss, Handler.CreateEmployeeHandler)
	route.DELETE("/:id/delete", Handler.ImplementrightManagementBoss, Handler.ImplementrightManagementMaster, Handler.DeleteByIdHandler)
	route.GET("/get", Handler.ImplementrightManagementBoss, Handler.GetAllEmployeesHandler)
	route.PUT("/:id/update", Handler.ImplementrightManagementMaster, Handler.UpdateById)
}
