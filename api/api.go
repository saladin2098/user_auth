package api

import (
	"github.com/Mubinabd/auth_service/api/handlers"
	"github.com/Mubinabd/auth_service/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Mubinabd/auth_service/docs"
)

// @Title UserAuth group swagger UI
// @BasePath /user
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:7070"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsConfig))

	r.Use(middleware.Middleware())

	user := r.Group("/user")
	{
		user.POST("/login", h.LoginUser)
		user.POST("/register", h.RegisterUser)
		user.GET("/info/:username", h.GetUserInfo)
	}
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
