package models

import "gorm.io/gorm"

type Authority struct {
	gorm.Model
	Authority     string
	AuthorityName string
}
