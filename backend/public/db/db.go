package db

import (
	"context"

	"github.com/blastertwist/antex-dash/config"
	"github.com/blastertwist/antex-dash/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func ConnectDB(cfg *config.Config, logger logger.Logger) *sqlx.DB {

	ctx := context.Background()

	dbSource := "postgresql://" + cfg.DB.User + ":" + cfg.DB.Pass + "@" + cfg.DB.URL + ":" + cfg.DB.Port + "/" + cfg.DB.Name + "?sslmode=disable"
	db, err := sqlx.ConnectContext(ctx, "postgres", dbSource)

	if err != nil {
		logger.Panic("[DB]Failed to connect to database.", zap.String("Error:", err.Error()))
	}

	err = db.PingContext(ctx)

	if err != nil {
		logger.Panic("[DB]Failed to connect to database.", zap.String("Error:", err.Error()))
	}

	return db

}
