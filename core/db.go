package core

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //
)

// OpenDb opens connection to database based on config
func OpenDb() *gorm.DB {
	config := ReadConfig()
	connection := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.SslMode,
	)
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	return db
}
