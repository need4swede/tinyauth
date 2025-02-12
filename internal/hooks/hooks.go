package hooks

import (
	"tinyauth/internal/auth"
	"tinyauth/internal/providers"
	"tinyauth/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func NewHooks(auth *auth.Auth, providers *providers.Providers) *Hooks {
	return &Hooks{
		Auth:      auth,
		Providers: providers,
	}
}

type Hooks struct {
	Auth      *auth.Auth
	Providers *providers.Providers
}

func (hooks *Hooks) UseUserContext(c *gin.Context) types.UserContext {
	cookie, cookiErr := hooks.Auth.GetSessionCookie(c)

	if cookiErr != nil {
		log.Error().Err(cookiErr).Msg("Failed to get session cookie")
		return types.UserContext{
			Username:   "",
			IsLoggedIn: false,
			OAuth:      false,
			Provider:   "",
		}
	}

	if cookie.Provider == "username" {
		log.Debug().Msg("Provider is username")
		if hooks.Auth.GetUser(cookie.Username) != nil {
			log.Debug().Msg("User exists")
			return types.UserContext{
				Username:   cookie.Username,
				IsLoggedIn: true,
				OAuth:      false,
				Provider:   "",
			}
		}
	}

	log.Debug().Msg("Provider is not username")
	provider := hooks.Providers.GetProvider(cookie.Provider)

	if provider != nil {
		log.Debug().Msg("Provider exists")
		if !hooks.Auth.EmailWhitelisted(cookie.Username) {
			log.Error().Str("email", cookie.Username).Msg("Email is not whitelisted")
			hooks.Auth.DeleteSessionCookie(c)
			return types.UserContext{
				Username:   "",
				IsLoggedIn: false,
				OAuth:      false,
				Provider:   "",
			}
		}
		log.Debug().Msg("Email is whitelisted")
		return types.UserContext{
			Username:   cookie.Username,
			IsLoggedIn: true,
			OAuth:      true,
			Provider:   cookie.Provider,
		}
	}

	return types.UserContext{
		Username:   "",
		IsLoggedIn: false,
		OAuth:      false,
		Provider:   "",
	}
}
