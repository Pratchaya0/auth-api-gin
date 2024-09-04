package utils

import (
	"errors"
	"fmt"

	"github.com/Pratchaya0/auth-api-gin/configs"
)

// Function สำหรับต่อ URL ในส่วนของ Database และ App ใน Utils
func ConnectionUrlBuilder(stuff string, cfg *configs.Configs) (string, error) {
	var dsn string

	switch stuff {
	case "fiber":
	case "gin":
		dsn = fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)
	case "postgresql":
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.PostgreSQL.Host,
			cfg.PostgreSQL.Port,
			cfg.PostgreSQL.Username,
			cfg.PostgreSQL.Password,
			cfg.PostgreSQL.Database,
			cfg.PostgreSQL.SSLMode,
		)
	default:
		errMsg := fmt.Sprintf("error, connection url builder doesn't know the %s", stuff)
		return "", errors.New(errMsg)
	}

	return dsn, nil
}
