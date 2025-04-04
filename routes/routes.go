package routes

import (
	"github.com/VoDangCMU/CMU-CS_445_LIS-LAB_4/api/auth"
	"github.com/VoDangCMU/CMU-CS_445_LIS-LAB_4/api/check"
	"github.com/VoDangCMU/CMU-CS_445_LIS-LAB_4/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	apiRoutes := r.Group("/api")
	{
		userRoutes := apiRoutes.Group("/user")
		{
			authRotues := userRoutes.Group("/auth")
			{
				authRotues.POST("/login", auth.Authentication)
				authRotues.PUT("/register", auth.Register)
				authRotues.POST("/logout", auth.Logout, middlewares.AuthMiddleware())
			}
			userRoutes.GET("/check-with-token", middlewares.AuthMiddleware(), check.CheckWithToken)
		}
	}
	return r
}
