package manager

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jutionck/golang-db-sinar-harapan-makmur/config"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type InfraManager interface {
	Conn() *sqlx.DB
	Log() *logrus.Logger
	LogFilePath() string
}

type infraManager struct {
	db  *sqlx.DB
	cfg *config.Config
	log *logrus.Logger
}

func (i *infraManager) LogFilePath() string {
	return i.cfg.LogFilePath
}

func (i *infraManager) Log() *logrus.Logger {
	logger := logrus.New()
	return logger
}

func (i *infraManager) Conn() *sqlx.DB {
	return i.db
}

func (i *infraManager) initDb() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		i.cfg.Host,
		i.cfg.Port,
		i.cfg.User,
		i.cfg.Password,
		i.cfg.Name,
	)
	conn, err := sqlx.Connect(i.cfg.Driver, dsn)
	if err != nil {
		panic(err)
	}
	i.db = conn
	return nil
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{cfg: cfg}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
