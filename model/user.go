package model

import (
	"strings"
	"time"
)

type User struct {
	Seq           int
	Email         string
	Name          string
	Password      string
	ProfileImgUrl string
	Bio           string
	IdProvider    IdentityProvider
	RegisterDate  time.Time
	UpdateDate    time.Time
}

type IdentityProvider int

const (
	Facebook IdentityProvider = 1 + iota
	GooglePlus
	Twitter
	Self
)

var IdentityProviders = [...]string{
	"Facebook",
	"Google+",
	"Twitter",
	"Self",
}

func (idp IdentityProvider) String() string {
	return IdentityProviders[idp-1]
}

func GetIdentityProvider(idp string) IdentityProvider {
	switch strings.ToLower(idp) {
	case "facebook":
		return Facebook
	case "google+":
		return GooglePlus
	case "twitter":
		return Twitter
	}
	return Self
}
