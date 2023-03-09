package config

import (
	"fmt"
	"os"
)

type Config struct {
	PG
}

type PG struct {
	PG_URL        string
	MAX_POOL_SIZE string
}

func New() *Config {
	cfg := &Config{
		PG: PG{},
	}

	pgUrl := os.Getenv("pg_url")
	fmt.Println(pgUrl)

	cfg.PG_URL = pgUrl

	return cfg
}
