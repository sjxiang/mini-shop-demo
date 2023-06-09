package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/sjxiang/mini-shop-demo/go-grpc-auth-svc/pkg/model"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,  // 不允许外键
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	db.AutoMigrate(new(model.User))

	return Handler{DB: db}
}
