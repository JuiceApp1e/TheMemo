package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wxcloudrun-golang/common"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/service"
)

func CreateTODO(c *gin.Context) {
	todo := &model.TodoModel{}
	c.BindJSON(todo)
	if err := service.SaveTodo(todo); err == nil {
		c.JSON(http.StatusOK, common.Ok("新增成功"))
	} else {
		c.JSON(http.StatusOK, common.Error("新增失败", err.Error()))
	}
}

func GetAllTodos(c *gin.Context) {
	if list, err := service.GetTodoList(); err == nil {
		c.JSON(http.StatusOK, common.Ok(list))
	} else {
		c.JSON(http.StatusOK, common.Error("查询失败", err.Error()))
	}
}

func UpdateTodo(c *gin.Context) {
	todo := &model.TodoModel{}
	c.BindJSON(todo)
	if err := service.UpdateTodo(todo); err == nil {
		c.JSON(http.StatusOK, common.Ok("更新成功"))
	} else {
		c.JSON(http.StatusOK, common.Error("更新失败", err.Error()))
	}
}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	idd, _ := strconv.ParseInt(id, 10, 32)
	if err := service.DeleteTodo(int32(idd)); err == nil {
		c.JSON(http.StatusOK, common.Ok("删除成功"))
	} else {
		c.JSON(http.StatusOK, common.Error("删除失败", err.Error()))
	}
}
