package config

import (
	"os"
)

func GetApplicationKey() string {
	return os.Getenv("SMARTY_SEED_APPLICATION_KEY")
}
