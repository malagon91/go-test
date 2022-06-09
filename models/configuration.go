package models

// Configuration :
type Configuration struct {
	Port               string   `json:"port" envconfig:"PORT"`
	CorsAllowedOrigins []string `json:"corsAllowedOrigins"`
	CorsAllowedMethods []string `json:"corsAllowedMethods"`
	CorsAllowedHeaders []string `json:"corsAllowedHeaders"`
	Environment        string   `envconfig:"GO_ENV"`
	LogLevel           string   `envconfig:"LOG_LEVEL"`
	Version            string   `envconfig:"VERSION"`
}
