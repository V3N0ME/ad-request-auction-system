package database

import (
	"database/sql"
	"fmt"
	models "janusapi/pkg/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//GetMysqlConnection returns a mysql connectiom with the input config
func GetMysqlConnection(config *models.MysqlConfig) (*sql.DB, error) {

	connectionURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Database)
	conn, err := sql.Open("mysql", connectionURL)

	if err != nil {
		return nil, err
	}

	if conn.Ping() != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(0)
	conn.SetConnMaxLifetime(time.Second * 10)

	return conn, nil
}
