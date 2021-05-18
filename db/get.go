package db

import "gorm.io/gorm"

// Get returns db connection
func Get() *gorm.DB {
	return databaseConn
}
