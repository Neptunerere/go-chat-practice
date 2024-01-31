package utils

import (
	"gorm.io/driver/mysql"
	"log"
	"time"

	"gorm.io/gorm"
)

var Connector *gorm.DB

func Connect(configStr string) (*gorm.DB, error) {

	var err error
	var connector *gorm.DB

	/// MySQL Connect Loop in Success
	for i := 0; i < 10; i++ {
		connector, err = gorm.Open(mysql.Open(configStr), &gorm.Config{})
		if err != nil {
			log.Printf("Unable to Open DB: %s... Retrying\n", err.Error())
			time.Sleep(time.Second * 2)

		}
		db, err1 := connector.DB()
		if err := db.Ping(); err != nil || err1 != nil {
			log.Printf("Unable to Ping DB: %s... Retrying\n", err.Error())
			time.Sleep(time.Second * 2)
		} else {
			err = nil
			break
		}

	}

	log.Println("Conn Successes")
	return connector, nil
}

// Migrate create/updates database table
func Migrate(table *interface{}) {
	err := Connector.AutoMigrate(&table)
	if err != nil {
		return
	}

	log.Println("Table migrated")
}
