package unionbank

import (
	"fmt"
	"github.com/jmramos02/smarty-seed-backend/config"
	net "net/url"
)

type GenerateUnionBankURLRequest struct {
	ProjectID uint `json:"project_id" validate:"required"`
	Amount    int  `json:"amount" validate:"required"`
	UserID    uint `json:"user_id"`
}

func GenerateUnionbankString(r GenerateUnionBankURLRequest, state string) string {
	url := fmt.Sprintf("%v/customers/v1/oauth2/authorize?client_id=%v&state=%v&partner_id=%v&response_type=code&scope=payments&type=single&redirect_uri=%v", config.GetUnionbankUrl(), net.QueryEscape(config.GetUnionBankClientID()), net.QueryEscape(state), net.QueryEscape(config.GetPartnerID()), net.QueryEscape(config.GetRedirectUri()))

	return url
}
