package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type Setting struct {
	Host string
	Port int
	User string
	Password string
	DBName string
}

const driverName = "mysql"

func InitDB(s Setting) *sql.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		s.User, s.Password, s.Host, s.Port, s.DBName)

	db, err := sql.Open(driverName, source)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := db.PingContext(ctx); err != nil {
		logrus.Fatal(err)
	}

	return db
}

func NewSetting(
	host string,
	port int,
	user, password, dbname string,
	) Setting {
	return Setting{
		Host: host,
		Port: port,
		User: user,
		Password: password,
		DBName: dbname,
	}
}
