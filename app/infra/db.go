package infra

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/api-sample/app/pkg/consts"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/golang-migrate/migrate/v4"
	m_mysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	host     = os.Getenv("DB_HOST")
	user     = os.Getenv("DB_USER")
	pwd      = os.Getenv("DB_PWD")
	database = os.Getenv("DB_DATABASE")
	port     = os.Getenv("DB_PORT")
	params   = "parseTime=true"
	DB       *gorm.DB
	err      error
	mySQLDB  *sql.DB
	Migrate  *migrate.Migrate
	Fixture  *testfixtures.Loader
)

func MysqlInit(includeDatabaseName bool) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", user, pwd, host, port, database, params)

	if !includeDatabaseName {
		dsn = fmt.Sprintf("%s:%s@(%s:%s)/?%s", user, pwd, host, port, params)
	}

	mySQLDB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
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
		MysqlInit(includeDatabaseName)
		return mySQLDB
	}

	return mySQLDB
}

func TestDatabaseInit(fixtreDir string) {
	if !consts.IsTest() {
		panic("Only the test environment can be run")
	}
	database = fmt.Sprintf("%s_test", database)
	MysqlInit(false)
	dropTestDatabase()
	DB.Exec("CREATE SCHEMA " + database)
	MysqlInit(true)
	driver, err := m_mysql.WithInstance(mySQLDB, &m_mysql.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/migrations", fixtreDir),
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", user, pwd, host, port, database, params)
	db, _ := sql.Open("mysql", dsn)
	f, _ := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(fmt.Sprintf("%s/fixtures", fixtreDir)),
	)
	Migrate = m
	Fixture = f
}

func dropTestDatabase() {
	if !consts.IsTest() {
		panic("Only the test environment can be run")
	}
	type Result struct {
		Database string
	}
	result := []Result{}
	DB.Raw("SHOW DATABASES").Scan(&result)
	for _, r := range result {
		if strings.Contains(r.Database, "test") {
			DB.Exec("DROP SCHEMA IF EXISTS " + r.Database)
		}
	}
}
