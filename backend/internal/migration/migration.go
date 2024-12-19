package migration

import (
	"Bot-or-Not/internal/domain/entity"
	"Bot-or-Not/internal/infra/database"
	"Bot-or-Not/pkg/config"
	"log/slog"
)

func init() {
	config.LoadEnv()

	db := database.New()
	defer func() {
		if err := db.Close(); err != nil {
			slog.Error("Failed to close the database connection", "error", err)
		} else {
			slog.Info("Database connection closed successfully")
		}
	}()

	if err := db.Conn().AutoMigrate(&entity.Player{}); err != nil {
		slog.Error("Database migration failed", "error", err)
		return
	}
	slog.Info("Database migration completed successfully")

}
