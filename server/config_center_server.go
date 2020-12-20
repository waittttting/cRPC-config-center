package server

import (
	"github.com/waittttting/cRPC-config-center/conf"
)

type ConfigCenterServer struct {
	config *conf.CCSConf
}

func NewConfigCenterServer(config *conf.CCSConf) *ConfigCenterServer {
	return &ConfigCenterServer{config: config}
}

func (ccs *ConfigCenterServer) Start() {

}