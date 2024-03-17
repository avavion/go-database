package config

import "github.com/joho/godotenv"

type (
	Config struct {
		Database Database
	}

	Database struct {
		PoolMax int
		URL     string
	}
)

func NewConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	cfg.Database = Database{
		PoolMax: 10,
		URL:     "root@/go_lang_database",
	}

	return cfg, nil
}
