package conf

type CCSConf struct {
	Server Server
	MySQL MySQL
}

type Server struct {
	Port int
}

type MySQL struct {
	Host string
	User string
	Password string
}