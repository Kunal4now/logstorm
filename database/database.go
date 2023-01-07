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
var DB *gorm.DB
	
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
	DB = db

	return nil
}

func CreateLog(level string, message string, timestamp time.Time, tag string, data string) (model.Log, error) {
	var newLog = model.Log{Level: level,  Message: message, Timestamp: timestamp, Tag: tag, Data: data}

	res := DB.Create(&model.Log{Level: level,  Message: message, Timestamp: timestamp, Tag: tag, Data: data})

	if res.Error != nil {
		return newLog, res.Error
	}

	return newLog, nil
}

func GetLogs(level string, tag string) ([]model.Log, error) {
	var logs []model.Log

	query := DB.Model(&model.Log{})

	if level != "" {
		query = query.Where("level = ?", level)
	}

	if tag != "" {
		query = query.Where("tag = ?", tag)
	}

	res := query.Find(&logs)

	if res.Error != nil {
		return nil, res.Error
	}

	return logs, nil
}

func GetLog(id string) (model.Log, error) {
	var log model.Log

	res := DB.Where("id = ?", id).First(&log)

	if res.Error != nil {
		return model.Log{}, res.Error
	}

	return log, nil
}

func DeleteLog(id string) error {
	res := DB.Where("id = ?", id).Delete(&model.Log{})

	if res.Error != nil {
		return res.Error
	}

	return nil
}
