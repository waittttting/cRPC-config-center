package server

import (
	"github.com/gin-gonic/gin"
	common "github.com/waittttting/cRPC-common"
	http2 "github.com/waittttting/cRPC-common/http"
	"github.com/waittttting/cRPC-common/model"
	"gorm.io/gorm"
	"net/http"
)

type WebHandler struct {
	db *gorm.DB
}

func NewWebHandler(db *gorm.DB) *WebHandler {
	return &WebHandler{
		db: db,
	}
}

func (wh *WebHandler) GetConfig(c *gin.Context) {

	serverName := c.PostForm("server_name")
	serverVersion := c.PostForm("server_version")

	sc := model.ServerConfig{
		ServerName: serverName,
		ServerVersion: serverVersion,
	}
	result := wh.db.Table("server_config").Where(&sc).Find(&sc)
	if result.Error != nil {
		c.JSON(http.StatusOK, http2.NewJSONResponseErr(common.ErrDB.ErrCode, common.CX_FAIL, common.ErrDB.ErrMsg, nil))
	}
	c.JSON(http.StatusOK, sc)
}


func (wh *WebHandler) SetConfig(c *gin.Context) {

}