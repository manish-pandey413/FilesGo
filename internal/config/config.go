package config

import "net/http"

type Config struct {
	App *Application
}

func New(serveMux *http.ServeMux) *Config {
	return &Config{
		App: NewApplication(serveMux),
	}
}
