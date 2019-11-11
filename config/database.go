package config

import (
	"os"
)

func GetDatabaseUrl() string {
	return os.Getenv("SMARTY_SEED_DB_URL")
}
