package models

import "gorm.io/gorm"

type Provider struct {
	gorm.Model
	ProviderID   string
	ProviderName string
	Enabled      string
}
