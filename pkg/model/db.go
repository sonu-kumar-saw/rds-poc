package model

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type databaseConfig struct {
	Host     string // RDS instance hostname or endpoint
	Port     int    // Port number for the database
	User     string // Database username
	Password string // Database password
	Name     string // Database name
}

type Database struct {
	logger    *zap.SugaredLogger
	mysqlConn *gorm.DB
}

func (dbConf *databaseConfig) init() {
	dbConf.Host = "localhost"
	dbConf.Port = 3306
	dbConf.User = "root"
	dbConf.Password = "root"
	dbConf.Name = "servo"

}

func (db *Database) connect(conf *databaseConfig) {
	//user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True", conf.User, conf.Password, conf.Host, conf.Port, conf.Name)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		db.logger.Fatalln("error occured while making connection to the database")
		return
	}
	db.mysqlConn = conn
}

func (db *Database) Init(logger *zap.SugaredLogger) {
	db.logger = logger
	conf := databaseConfig{}
	conf.init()
	db.logger.Infow("attempting for mysql db connection",
		"conf", conf)
	db.connect(&conf)

	// create all the tables if already not existing.
	db.mysqlConn.AutoMigrate(
		&Ingest{}, // Auto create table named Ingest if not exisiting.
	)
}
