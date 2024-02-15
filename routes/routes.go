package routes

import (
	atom_user "user_service/atom/user/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	route := gin.Default()

	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	user := route.Group("user")
	{
		user.GET("get/all", atom_user.GetAllData)
		user.GET("get/:id", atom_user.GetDataUserById)
		user.POST("get/name", atom_user.GetDataUserByName)
		user.POST("add", atom_user.CreateUser)
		user.POST("update/:id", atom_user.UpdateUser)
	}

	return route
}
