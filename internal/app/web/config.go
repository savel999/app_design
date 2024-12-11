package web

type Config struct {
	ServiceName string
	LogLevel    string
	ServerAddr  string
}

func InitConfig() *Config {
	return &Config{
		ServiceName: "web",
		LogLevel:    "info",
		ServerAddr:  ":8080",
	}
}
