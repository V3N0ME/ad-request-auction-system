package models

//MysqlConfig is the config model required to create a mysql connection
type MysqlConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Database string
}
