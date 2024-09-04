package providers

import (
	"github.com/Pratchaya0/auth-api-gin/configs"
	"github.com/Pratchaya0/auth-api-gin/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDBConnection(cfg *configs.Configs) (*gorm.DB, error) {
	dsn, err := utils.ConnectionUrlBuilder("postgresql", cfg)
	if err != nil {
		return nil, err
	}

	databases, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return databases, nil
}
