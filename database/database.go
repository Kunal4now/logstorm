package database

import (
	"time"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"github.com/Kunal4now/logstorm/model"
)

var dsn string
	
func InitDB() error {
	var (
		host = viper.GetString("HOST")
		port = viper.GetInt("PORT")
		user = viper.GetString("USER")
		password = viper.GetString("PASSWORD")
		dbname = viper.GetString("DBNAME")
	)

	dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Kolkata", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.AutoMigrate(&model.Log{})

	return nil
}

func CreateLog(level string, message string, timestamp time.Time, tag string, data string) (model.Log, error) {
	var newLog = model.Log{Level: level,  Message: message, Timestamp: timestamp, Tag: tag, Data: data}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return newLog, err
	}

	db.Create(&model.Log{Level: level,  Message: message, Timestamp: timestamp, Tag: tag, Data: data})

	return newLog, nil
}

func GetLogs(level string, tag string) ([]model.Log, error) {
	var logs []model.Log

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return logs, err
	}

	query := db.Model(&model.Log{})

	if level != "" {
		query = query.Where("level = ?", level)
	}

	if tag != "" {
		query = query.Where("tag = ?", tag)
	}

	query.Find(&logs)

	return logs, nil
}

func GetLog(id string) (model.Log, error) {
	var log model.Log

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return log, err
	}

	db.Where("id = ?", id).First(&log)

	return log, nil
}

func DeleteLog(id string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Where("id = ?", id).Delete(&model.Log{})

	return nil
}
