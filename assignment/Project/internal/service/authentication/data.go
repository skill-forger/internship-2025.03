package authentication

import (
	"github.com/spf13/viper"

	ct "golang-project/internal/contract"
	m "golang-project/internal/model"
	"golang-project/static"
)

// prepareSignInResponse transforms the data and returns the Sign In Response
func prepareSignInResponse(o *m.User, token string) *ct.SignInResponse {
	return &ct.SignInResponse{
		UserID:       o.ID,
		Token:        token,
		Type:         viper.GetString(static.EnvAuthType),
		ExpiredAfter: viper.GetInt(static.EnvAuthLifeTime),
	}
}
