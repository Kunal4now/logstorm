package database

import (
	"time"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

type Log struct {
	gorm.Model
	Level string `gorm:"type:varchar(255)"`
	Message string `gorm:"type:varchar(255)"`
	Timestamp time.Time `gorm:"type:timestamp"`
	Tag string `gorm:"type:varchar(255)"`
	Data struct{} `gorm:"type:jsonb"`
}

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

	db.AutoMigrate(&Log{})

	return nil
}

func CreateLog(level string, message string, timestamp time.Time, tag string, data struct{}) (Log, error) {
	var newLog = Log{Level: level,  Message: message, Timestamp: timestamp, Tag: tag, Data: data}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return newLog, err
	}

	db.Create(&Log{Level: level,  Message: message, Timestamp: timestamp, Tag: tag, Data: data})

	return newLog, nil
}

func GetLogs() ([]Log, error) {
	var logs []Log

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return logs, err
	}

	db.Find(&logs)

	return logs, nil
}

func GetLog(id string) (Log, error) {
	var Log Log

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return Log, err
	}

	db.Where("id = ?", id).First(&Log)

	return Log, nil
}

func DeleteLogs(id string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Where("id = ?", id).Delete(&Log{})

	return nil
}
