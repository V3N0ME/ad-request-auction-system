package errors

var (
	//MysqlBConnectionError is the error returned when mysql master db is down
	MysqlBConnectionError = New("Mysql Master db is down", 101)
)
