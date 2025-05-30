package quickwit

type Config struct {
	Endpoint string `mapstructure:"endpoint" env:"QUICKWIT_ENDPOINT"`
}
