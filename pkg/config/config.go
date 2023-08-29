package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	postgresUser    string
	postgresPass    string
	postgresHost    string
	postgresPort    string
	postgresDB      string
	postgresSSLMode string
	env             string
	httpPort        string
	httpHost        string
	historyDir      string
}

func (cfg config) GetPostgresUser() string {
	return cfg.postgresUser
}

func (cfg config) GetPostgresPass() string {
	return cfg.postgresPass
}

func (cfg config) GetPostgresHost() string {
	return cfg.postgresHost
}

func (cfg config) GetPostgresPort() string {
	return cfg.postgresPort
}

func (cfg config) GetPostgresDB() string {
	return cfg.postgresDB
}

func (cfg config) GetPostgresSSLMode() string {
	return cfg.postgresSSLMode
}

func (cfg config) GetEnv() string {
	return cfg.env
}

func (cfg config) GetHttpPort() string {
	return cfg.httpPort
}

func (cfg config) GetHttpURL() string {
	return cfg.httpHost + ":" + cfg.httpPort
}
func (cfg config) GetHistoryDir() string {
	return cfg.historyDir
}

func LoadEnv(filenames ...string) error {
	const op = "pkg.config.LoadEnv"
	err := godotenv.Load(filenames...)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func GetConfig() *config {
	cfg := &config{
		postgresUser:    "",
		postgresPass:    "",
		postgresHost:    "localhost",
		postgresPort:    "27017",
		postgresDB:      "",
		env:             "local",
		postgresSSLMode: "disable",
		httpHost:        "localhost",
		historyDir:      "history",
	}

	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DB")
	ssl := os.Getenv("POSTGRES_SSL_MODE")
	env := os.Getenv("ENV")
	httpPort := os.Getenv("HTTP_PORT")
	httpHost := os.Getenv("HTTP_HOST")

	if env != "" {
		cfg.env = env
	}
	if httpPort != "" {
		cfg.httpPort = httpPort
	}
	if httpHost != "" {
		cfg.httpHost = httpHost
	}
	if user != "" {
		cfg.postgresUser = user
	}
	if pass != "" {
		cfg.postgresPass = pass
	}
	if host != "" {
		cfg.postgresHost = host
	}
	if port != "" {
		cfg.postgresPort = port
	}
	if db != "" {
		cfg.postgresDB = db
	}
	if ssl != "" {
		cfg.postgresSSLMode = ssl
	}

	return cfg
}
