package build

import "gorm.io/gorm"

type DatabaseManager struct {
	*gorm.DB
}

func NewDatabaseManager(db *gorm.DB) *DatabaseManager {
	return &DatabaseManager{DB: db}
}

type Controller struct {
	DBManager *DatabaseManager
}

func NewController(dbManager *DatabaseManager) *Controller {
	return &Controller{DBManager: dbManager}
}
