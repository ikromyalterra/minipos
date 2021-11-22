package config

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	//driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	//DB DB
	DB *gorm.DB
)

//Database Database
type Database struct {
	Host      string
	User      string
	Password  string
	DBName    string
	DBNumber  int
	Port      int
	DebugMode bool
}

// LoadDBConfig load database configuration
func LoadDBConfig(name string) Database {
	db := viper.Sub("database." + name)
	conf := Database{
		Host:      db.GetString("host"),
		User:      db.GetString("user"),
		Password:  db.GetString("password"),
		DBName:    db.GetString("db_name"),
		Port:      db.GetInt("port"),
		DebugMode: db.GetBool("debug"),
	}
	return conf
}

func OpenDbPool() {
	DB = MysqlConnect("mysql")
	pool := viper.Sub("database.mysql.pool")
	DB.DB().SetMaxOpenConns(pool.GetInt("maxOpenConns"))
	DB.DB().SetMaxIdleConns(pool.GetInt("maxIdleConns"))
	DB.DB().SetConnMaxLifetime(pool.GetDuration("maxLifetime") * time.Second)
}

// MysqlConnect connect to mysql using config name. return *gorm.DB incstance
func MysqlConnect(configName string) *gorm.DB {
	mysql := LoadDBConfig(configName)
	connection, err := gorm.Open("mysql", mysql.User+":"+mysql.Password+"@tcp("+mysql.Host+":"+strconv.Itoa(mysql.Port)+")/"+mysql.DBName+"?charset=utf8&parseTime=True&loc=Local")
	// enable debug
	if err != nil {
		panic(err)
	}

	connection.LogMode(true)

	if mysql.DebugMode {
		return connection.Debug()
	}

	return connection
}
