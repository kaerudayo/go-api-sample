package infra

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/api-sample/app/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	host      = os.Getenv("DB_HOST")
	user      = os.Getenv("DB_USER")
	pwd       = os.Getenv("DB_PWD")
	database  = os.Getenv("DB_DATABASE")
	port      = os.Getenv("DB_PORT")
	dnsParams = "charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local&multiStatements=true"
	DB        *gorm.DB
	err       error
	mySQLDB   *sql.DB
)

func Init(includeDatabaseName bool) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", user, pwd, host, port, database, dnsParams)

	if !includeDatabaseName {
		dsn = fmt.Sprintf("%s:%s@(%s:%s)/?%s", user, pwd, host, port, dnsParams)
	}

	mySQLDB, err = sql.Open("mysql", dsn)
	if err != nil {
		logger.SugerError(err.Error())
	}

	mySQLDB.SetMaxOpenConns(20)
	mySQLDB.SetMaxIdleConns(10)
	mySQLDB.SetConnMaxLifetime(time.Hour)

	DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: mySQLDB,
	}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	if err != nil {
		fmt.Printf("db open error: dsn: %s, error: %s", dsn, err.Error())
		time.Sleep(5 * time.Second)
		Init(includeDatabaseName)
		return mySQLDB
	}

	return mySQLDB
}
