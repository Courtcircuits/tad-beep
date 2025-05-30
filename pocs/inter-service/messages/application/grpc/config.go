package grpc

type Config struct {
	ListenAddr string `mapstructure:"listen_addr" env:"GRPC_LISTEN_ADDR"`
}
