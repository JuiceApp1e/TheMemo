package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var mysqlDbInstance *gorm.DB
var mongoDbInstance *mongo.Database

// Init 初始化数据库
func Init() error {
	if err := MysqlInit(); err != nil {
		return err
	}
	if err := MongoDbInit(); err != nil {
		return err
	}
	return nil
}

func MysqlInit() error {
	source := "%s:%s@tcp(%s)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true"
	user := os.Getenv("MYSQL_USERNAME")
	pwd := os.Getenv("MYSQL_PASSWORD")
	addr := os.Getenv("MYSQL_ADDRESS")
	dataBase := os.Getenv("MYSQL_DATABASE")
	if dataBase == "" {
		dataBase = "golang_demo"
	}
	source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	log.Println("start init mysql with ", source)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		}})
	if err != nil {
		log.Fatalln("DB Open error,err=", err.Error())
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("DB Init error,err=", err.Error())
		return err
	}

	// 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(100)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(200)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	mysqlDbInstance = db

	log.Println("finish init mysql with ", source)
	return nil
}

func MongoDbInit() error {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://JuiceApple:guodie6666666@39.103.202.134:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	log.Println("Connected to MongoDB!")
	// 指定获取要操作的数据集
	mongoDb := client.Database("memo")

	mongoDbInstance = mongoDb

	return nil
}

// GetMysqlDb ...
func GetMysqlDb() *gorm.DB {
	return mysqlDbInstance
}

func GetMongoDb() *mongo.Database {
	return mongoDbInstance
}
