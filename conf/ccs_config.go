package conf

import "time"

type CCSConf struct {
	Server Server
	MySQL MySQL
}

type Server struct {
	TcpPort int
	HttpPort int
	ReceiveSocketChanLen int
	ReceiveSocketTimeoutMs time.Duration
}

type MySQL struct {
	Host string
	User string
	Password string
}