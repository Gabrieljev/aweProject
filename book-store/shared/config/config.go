package config

import (
	"github.com/pkg/errors"
)

type (
	EnvConfiguration struct {
		// - sql-config-start

		SqlConnectionUri string `envconfig:"SQL_CONNECTION_URI" default:"root:root@tcp(127.0.0.1:3306)/book_store?charset=utf8&parseTime=True&loc=Local"`

		SqlMaxIdleConnection int `envconfig:"SQL_MAX_IDLE_CONNECTION" default:"4"`

		SqlMaxOpenConnection int `envconfig:"SQL_MAX_OPEN_CONNECTION" default:"8"`

		SqlConnMaxLifetime int `envconfig:"SQL_CONN_MAX_LIFETIME" default:"300"`

		SqlLogMode bool `envconfig:"MYSQL_LOG_MODE" default:"false"`

		SqlGetAvailabilityTimeout string `envconfig:"SQL_GET_AVAILABILITY_TIMEOUT" default:"1s"`

		// - sql-config-end

		// - echo-config-start

		EchoServerPort int `envconfig:"ECHO_SERVER_PORT" default:"9001"`

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
