package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq"
	"io/ioutil"
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

func GetOpenId(c *gin.Context) {
	const (
		code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
		appID           = "wx897f70f528759690"
		appSecret       = "f3f76db79938c6e677b06810405fb16b"
	)
	//获取code
	code := c.Query("code")

	//调用auth.code2Session接口获取openid
	url := fmt.Sprintf(code2sessionURL, appID, appSecret, code)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json := gojsonq.New().FromString(string(body)).Find("openid")
	openId := json.(string)
	c.JSON(http.StatusOK, openId)
}
