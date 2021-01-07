package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/waittttting/cRPC-common/tcp"
	"github.com/waittttting/cRPC-config-center/conf"
	"github.com/waittttting/cRPC-config-center/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

func main() {

	var configPath string
	flag.StringVar(&configPath, "config", "", "config path")
	flag.Parse()

	var config conf.CCSConf
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		logrus.Fatalf("load local config err [%v]", err)
	}

	ccs := server.NewConfigCenterServer(&config)

	ccs.Start()
	// mysql
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/cx_config_center?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQL.User,
		config.MySQL.Password,
		config.MySQL.Host)

	configDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("open mysql err [%v]", err)
	}

	// tcp
	tcp.AcceptSocket(
		strconv.Itoa(config.Server.TcpPort),
		ccs.ReceiveSocketChan,
		config.Server.ReceiveSocketTimeoutMs)


	// web
	wh := server.NewWebHandler(configDb)
	r := gin.Default()
	get := r.Group("/get")
	get.POST("/config", wh.GetConfig)

	r.Run(fmt.Sprintf(":%v", config.Server.HttpPort))
}
