package store

type Config struct {
	DatabaseURL string `toml:"databaseurl"`
}

func NewConfig() *Config {
	return &Config{}
}
