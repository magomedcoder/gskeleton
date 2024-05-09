package config

type Http struct {
	Port int `yaml:"port"`
}

type Grpc struct {
	Host         string `yaml:"host"`
	GrpcProtocol string `yaml:"grpc_protocol"`
	GrpcPort     int    `yaml:"grpc_port"`
}

type Server struct {
	Http *Http `yaml:"http"`
	Grpc *Grpc `yaml:"grpc"`
}
