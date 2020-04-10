package dao

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type dbConfig struct {
	Mysql   mysqlConfig
	Mongodb mongodbConfig
}

type mysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Charset  string
}

func (c *mysqlConfig) toURI() string {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)
	return uri
}

type mongodbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	AuthDB   string
}

func (c *mongodbConfig) toURI() string {
	uri := fmt.Sprintf("mongodb://%s:%d/?connect=direct",
		c.Host, c.Port)
	return uri
}

var MysqlClient *sqlx.DB
var MongoClient *mongo.Client

func InitDatabase(mode string) {
	var err error

	// Read toml config file to get database config.
	var configFile string
	if mode == "release" {
		configFile = "configs/db_product.toml"
	} else {
		configFile = "configs/db_develop.toml"
	}

	var config dbConfig
	if _, err = toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}

	// Connect to MySQL
	if MysqlClient, err = sqlx.Open("mysql", config.Mysql.toURI()); err != nil {
		log.Fatal(err)
	}
	if err = MysqlClient.Ping(); err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB
	auth := options.Credential{
		AuthSource: config.Mongodb.AuthDB,
		Username:   config.Mongodb.User,
		Password:   config.Mongodb.Password,
	}
	option := options.Client().ApplyURI(config.Mongodb.toURI()).SetAuth(auth).SetMaxPoolSize(100)
	if MongoClient, err = mongo.Connect(context.TODO(), option); err != nil {
		log.Fatal(err)
	}
	if err = MongoClient.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
}
