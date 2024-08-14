package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName             string
	LastName              string
	UserName              string
	Password              []byte
	ProfilePicture        string
	Email                 string
	AccountNonExpired     string
	AccountNonLock        string
	CredentialsNonExpired string

	Provider    Provider
	Authorities []Authority
}
