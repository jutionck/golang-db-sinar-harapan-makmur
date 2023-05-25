package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DbConnection interface {
	Conn() *sqlx.DB
}

type dbConnection struct {
	db  *sqlx.DB
	cfg *Config
}

func (d *dbConnection) initDb() error {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.cfg.Host, d.cfg.Port, d.cfg.User, d.cfg.Password, d.cfg.Name)
	db, err := sqlx.Connect(d.cfg.Driver, psqlconn)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	d.db = db
	return nil
}

func (d *dbConnection) Conn() *sqlx.DB {
	return d.db
}

func NewDbConnection(cfg *Config) (DbConnection, error) {
	conn := &dbConnection{
		cfg: cfg,
	}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
