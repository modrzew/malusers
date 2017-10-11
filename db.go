package malusers

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
		config.Host,
		config.Port,
		config.Database,
		config.Username,
		config.Password,
		config.SslMode,
	)
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	return db
}
