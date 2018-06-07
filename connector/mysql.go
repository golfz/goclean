package connector

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"fmt"
)

func InitiateMysqlConnectionSession() *sql.DB {
	dbObj := initiateDbObject()
	session, _ := sql.Open("mysql", dbObj.getDataSourceName())
	return session
}

type dbObject struct {
	defaultData map[string]string
}

func initiateDbObject() *dbObject {
	data := map[string]string{
		"THE_DB_HOST":     "localhost",
		"THE_DB_PORT":     "3306",
		"THE_DB_DBNAME":   "EXAMPLE_DB",
		"THE_DB_USER":     "EXAMPLE",
		"THE_DB_PASSWORD": "1234567890",
	}
	return &dbObject{defaultData: data}
}

func (dbo *dbObject) get(key string) string {
	value := os.Getenv(key)

	if (value == "") {
		value = dbo.defaultData[key]
	}

	return value
}

func (dbo *dbObject) getDataSourceName() string {
	host := dbo.get("THE_DB_HOST")
	port := dbo.get("THE_DB_PORT")
	db := dbo.get("THE_DB_DBNAME")
	user := dbo.get("THE_DB_USER")
	pwd := dbo.get("THE_DB_PASSWORD")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pwd, host, port, db)
}
