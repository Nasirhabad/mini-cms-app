package routes

import "github.com/labstack/echo/v4"

func Setuproutes(e *echo.Echo) {
	// content routes
	contentController := controllers.Initcontentcontroller()

	contentRoutes := e.Group("/api/v1")

	contentRoutes.GET("/contents", contentController.Getall)
	contentRoutes.GET("/contents/:id", contentController.GetByID)
	contentRoutes.POST("/contents", contentController.Create)
	contentRoutes.PUT("/contents/:id", contentController.Update)
	contentRoutes.DELETE("/contents/:id", contentController.Delete)
}
