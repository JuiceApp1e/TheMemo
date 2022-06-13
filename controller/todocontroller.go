package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wxcloudrun-golang/service"
)

func SaveTodosList(c *gin.Context) {
	// 注意：下面为了举例子方便，暂时忽略了错误处理
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m map[string]interface{}
	// 反序列化
	_ = json.Unmarshal(b, &m)
	log.Println(m)

	openid := m["key"].(string)

	service.SaveTodos(openid, m["todos"])
	c.JSON(http.StatusOK, m)
}

func GetTodosList(c *gin.Context) {
	// 注意：下面为了举例子方便，暂时忽略了错误处理
	openid := c.Query("key")

	result, _ := service.GetTodos(openid)
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}
