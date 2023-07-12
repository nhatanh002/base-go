package config

import (
	"base-go/common/utils"

	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Env struct {
	Environment string
	HttpConfig  HttpConf
	DBconfig    DBconf
	Secrets     Secrets
}

func parseEnv() Env {

	// read env
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		viper.AutomaticEnv()
	}
	CORSstring := WithDefault("APP_CORS", "*")
	CORS := strings.Split(CORSstring, " ")
	fmt.Println(CORS)

	env := WithDefault("APP_ENV", LocalEnv)
	if !utils.Member(env, []string{LocalEnv, DevEnv, StagingEnv, ProductionEnv}) {
		env = LocalEnv
	}

	return Env{
		Environment: env,

		HttpConfig: HttpConf{
			Host:             WithDefault("APP_HOST", ""),
			Port:             WithDefault("APP_PORT", 8001),
			Mode:             WithDefault("APP_MODE", "debug"), // "release", "test"
			CORSAllowOrigins: CORS,

			X509CertFile: WithDefault("APP_X509_CERT", "./cert.pem"),
			X509KeyFile:  WithDefault("APP_X509_KEY", "./key.pem"),
		},

		DBconfig: DBconf{
			Host:            WithDefault("APP_DB_HOST", "localhost"),
			Port:            WithDefault("APP_DB_PORT", 3306),
			Username:        WithDefault("APP_DB_USERNAME", "dummy"),
			Password:        WithDefault("APP_DB_PASSWORD", ""),
			DBname:          WithDefault("APP_DB_NAME", "dummy"),
			DBbackend:       WithDefault("APP_DB_BACKEND", "mysql"),
			MaxOpenConns:    WithDefault("APP_DB_CONN_MAX", 10),
			MaxIdleConns:    WithDefault("APP_DB_CONN_MAX_IDLE", 5),
			ConnMaxLifetime: WithDefault("APP_DB_CONN_MAX_LIFETIME", 30*time.Second),
			SSLMode:         WithDefault("APP_DB_SSL_MODE", "disable"),
		},

		Secrets: Secrets{
			JwtSecret: WithDefault("APP_JWT_SECRET", "keepcalmandstaypositive"),
		},
	}
}

type Config struct {
	Env
}

var cnf *Config

func Load() {
	if cnf != nil {
		return
	}

	// parse env
	env := parseEnv()

	// init config
	cnf = &Config{
		Env: env,
	}
	return
}

func Get() *Config {
	if cnf == nil {
		Load()
	}
	return cnf
}
