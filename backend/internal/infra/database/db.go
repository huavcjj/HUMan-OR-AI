package database

import (
	"Bot-or-Not/pkg/config"
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB
}

func New() *DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort,
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	return &DB{conn: conn}
}

func (db *DB) Close() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from gorm.DB: %v", err)
	}
	return sqlDB.Close()
}

func (db *DB) WithContext(ctx context.Context) *gorm.DB {
	return db.conn.WithContext(ctx)
}

func (db *DB) Conn() *gorm.DB {
	return db.conn
}
