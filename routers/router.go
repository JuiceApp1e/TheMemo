package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wxcloudrun-golang/controller"
)

func SetupRouter() *gin.Engine {
	// 创建一个默认的路由引擎
	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})

	//v1 := r.Group("/v1")
	//{
	//	v1.POST("/todo", controller.CreateTODO)
	//	v1.DELETE("/todo/:id", controller.DeleteTodo)
	//	v1.PUT("/todo", controller.UpdateTodo)
	//	v1.GET("/todo", controller.GetAllTodos)
	//}

	v2 := r.Group("/v2")
	{
		v2.POST("/todo", controller.SaveTodosList)
		v2.GET("/todo", controller.GetTodosList)
		//v2.GET("/todo/:openid", controller.GetTodosIsInited)
	}
	return r
}
