package main

import (
	"database/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		panic("Неудалось загрузить конфиг")
	}

	db, err := sql.Open("mysql", cfg.Database.URL)

	if err != nil {
		panic("ошибка базы данных: " + err.Error())
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(cfg.Database.PoolMax)
	db.SetMaxIdleConns(cfg.Database.PoolMax)
}
