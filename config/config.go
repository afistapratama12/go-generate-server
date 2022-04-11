package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Environment       string `envconfig:"ENVIRONMENT" default:"production"`
	LocalGraderApiUrl string `envconfig:"LOCAL_GRADER_API_URL" default:"http://localhost:8080"`
}

// Get to get defined configuration
func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
