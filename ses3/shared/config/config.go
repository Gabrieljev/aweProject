package config

import (
	"github.com/pkg/errors"
)

type (
	EnvConfiguration struct {
		// - mongo-config-start

		MongoDatabaseUri string `envconfig:"MONGO_DATABASE_URI" default:"mongodb://0.0.0.0:27017"`

		MongoDatabaseName string `envconfig:"MONGO_DATABASE_NAME" default:"TIX-BOOK-STORE"`

		// - mongo-config-end

		// - echo-config-start

		EchoServerPort int `envconfig:"ECHO_SERVER_PORT" default:"9000"`

		// - end-of-config
	}
)

func NewEnvConfiguration() (*EnvConfiguration, error) {
	configuration := EnvConfiguration{}

	if err := NewFromEnv(&configuration); err != nil {
		return nil, errors.Wrap(err, "failed to provide env config")
	}

	return &configuration, nil
}
