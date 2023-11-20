package server

type ServerList struct {
	Grpc Server
}

// Server :
type Server struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	TLS     bool   `yaml:"tls"`
	Timeout int    `yaml:"timeout"`
}
