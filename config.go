package server

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type serverConfiguration struct {
	DB         *databaseConfiguration
	SecretSeed string `required:"true"`
}

type databaseConfiguration struct {
	Host     string `required:"true"`
	Port     string `default:"5432"`
	Name     string `required:"true"`
	Password string `required:"true"`
	User     string `required:"true"`
}

func ResolveConfig() (*serverConfiguration, error) {
	config := &serverConfiguration{}

	//
	// Resolve env. variables
	//
	if err := envconfig.Process("SERVER", config); err != nil {
		return nil, fmt.Errorf("failed to parse environment configurations, %s", err.Error())
	}

	return config, nil
}
