package server

import (
	"cx-rpc-config-center/conf"
)

type ConfigCenterServer struct {
	config *conf.CCSConf
}

func NewConfigCenterServer(config *conf.CCSConf) *ConfigCenterServer {
	return &ConfigCenterServer{config: config}
}

func (ccs *ConfigCenterServer) Start() {

}