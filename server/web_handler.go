package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WebHandler struct {

}

func NewWebHandler() *WebHandler {
	return &WebHandler{}
}

func (wh *WebHandler) GetConfig(c *gin.Context) {

	_ = c.PostForm("server_name")
	_ = c.PostForm("server_version")
	// todo: 没有正确读到配置信息，返回失败
	c.JSON(http.StatusOK, make(map[string]interface{}))
}