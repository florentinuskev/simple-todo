package db

import (
	"context"
	"log"

	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg *utils.Config) *sqlx.DB {

	ctx := context.Background()

	dbSource := "postgresql://" + cfg.Env["DB_USER"] + ":" + cfg.Env["DB_PASS"] + "@" + cfg.Env["DB_URL"] + ":" + cfg.Env["DB_PORT"] + "/" + cfg.Env["DB_NAME"] + "?sslmode=disable"
	db, err := sqlx.ConnectContext(ctx, "postgres", dbSource)

	if err != nil {
		log.Panic("[DB]Failed to connect to database.", err.Error())
	}

	err = db.PingContext(ctx)

	if err != nil {
		log.Panic("[DB]Failed to connect to database.", err.Error())
	}

	return db

}
