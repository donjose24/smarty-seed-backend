package config

import (
	"os"
)

//DON"T DO THIS HAHA
func GetUnionbankUrl() string {
	return os.Getenv("UNIONBANK_BASE_URL")
}

func GetUnionBankClientID() string {
	return os.Getenv("UNIONBANK_CLIENT_ID")
}

func GetUnionBankClientSecret() string {
	return os.Getenv("UNIONBANK_CLIENT_SECRET")
}

func GetPartnerID() string {
	return os.Getenv("UNIONBANK_PARTNER_ID")
}

func GetRedirectUri() string {
	return os.Getenv("UNIONBANK_REDIRECT_URI")
}
