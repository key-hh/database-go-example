package internal

type Config struct {
	Host              string
	Port              int
	ReadTimeout       int
	ReadHeaderTimeout int
}

func NewDefaultConfig() *Config {
	return &Config{
		Host:              "0.0.0.0",
		Port:              8090,
		ReadTimeout:       10,
		ReadHeaderTimeout: 3,
	}
}
