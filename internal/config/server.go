package config

type Grpc struct {
	Host         string `yaml:"host"`
	GrpcProtocol string `yaml:"grpc_protocol"`
	GrpcPort     string `yaml:"grpc_port"`
}

type JsonRpc struct {
	Port string `yaml:"port"`
}

type Server struct {
	JsonRpc *JsonRpc `yaml:"json_rpc"`
	Grpc    *Grpc    `yaml:"grpc"`
}
