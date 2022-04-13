package postgres

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/nashmaniac/hi-mama/hi-mama-backend/adapters"
	"github.com/nashmaniac/hi-mama/hi-mama-backend/models"
)

type postgresRepository struct {
	db *gorm.DB
}

func (r *postgresRepository) CloseDB(ctx context.Context) {
	sqlDb, _ := r.db.DB()
	log.Println("closing the database")
	sqlDb.Close()
}

func NewRepository(
	ctx context.Context,
	username string,
	password string,
	databaseName string,
	host string,
	port string,
) (adapters.PeristenceStore, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, username, password, databaseName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})
	log.Println("Connection to db is sucessful")
	if err != nil {
		return nil, err
	}

	// migration here
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Entry{})

	return &postgresRepository{
		db: db,
	}, nil
}
