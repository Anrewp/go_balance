package router

import (
	"net/http"

	"github.com/Anrewp/go_balance/src/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controller.UserController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	userRouter := router.Group("/users")
	userRouter.GET("", userController.FindAll)
	userRouter.GET("/:userId", userController.FindById)
	userRouter.POST("", userController.Create)
	userRouter.PATCH("/:userId", userController.Update)
	userRouter.DELETE("/:userId", userController.Delete)

	return service
}
