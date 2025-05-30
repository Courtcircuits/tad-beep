package grpc

type Config struct {
	Endpoint string `mapstructure:"endpoint" env:"GRPC_ENDPOINT"`
}

