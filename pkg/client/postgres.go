package client

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type configer interface {
	GetPostgresDB() string
	GetPostgresHost() string
	GetPostgresPass() string
	GetPostgresPort() string
	GetPostgresUser() string
	GetPostgresSSLMode() string
}

type postgresClient struct {
	host    string
	port    string
	user    string
	pass    string
	db      string
	sslMode string
}

func NewPostgresClient(cfg configer) *postgresClient {
	return &postgresClient{
		host: cfg.GetPostgresHost(),
		port: cfg.GetPostgresPort(),
		user: cfg.GetPostgresUser(),
		pass: cfg.GetPostgresPass(),
		db: cfg.GetPostgresDB(),
		sslMode: cfg.GetPostgresSSLMode(),
	}
}

func (cl *postgresClient) GetDb() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cl.user,
		cl.pass,
		cl.host,
		cl.port,
		cl.db,
		cl.sslMode,
	)
	return sqlx.Open("postgres", psqlInfo)
}
