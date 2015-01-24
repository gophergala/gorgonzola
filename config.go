package main

type Config struct {
	Server string
}

func NewConfig() *Config {
	return &Config{
		Server: ":8000",
	}
}
