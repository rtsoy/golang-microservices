package mysql

import (
	"users/internal/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(host.docker.internal:3306)/users"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(domain.User{})
}
