package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApi() {
	route := gin.Default()
	// To enable CORS
	route.Use(corsMiddleware())
	// Swagger Documentation
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// User Module Routes and Handlers
	setupUserModuleRoutes(route)

	err := route.Run()
	if err != nil {
		return
	}

}

// It holds User Module all Routes
func setupUserModuleRoutes(route *gin.Engine) {
	route.POST(Route_ChangePassword, changePassword)
	route.POST(Route_ForgotPassword, forgotPassword)
	route.POST(Route_IsUserExist, isUserExist)
	route.POST(Route_SignUp, signUp)
	route.POST(Route_SignIn, signIn)
}

func setupTeacherModuleRoutes(route *gin.Engine) {
	route.POST(Route_CreateTeacherProfile, createTeacherProfile)
	route.POST(Route_GetTeacherProfile, getTeacherProfile)
}

// To enable CORS request
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin) // Allow dynamic origins
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		}
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
