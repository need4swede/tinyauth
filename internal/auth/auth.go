package auth

import (
	"tinyauth/internal/types"

	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func NewAuth(userList types.Users, oauthWhitelist []string) *Auth {
	return &Auth{
		Users:          userList,
		OAuthWhitelist: oauthWhitelist,
	}
}

type Auth struct {
	Users          types.Users
	OAuthWhitelist []string
}

func (auth *Auth) GetUser(username string) *types.User {
	for _, user := range auth.Users {
		if user.Username == username {
			return &user
		}
	}
	return nil
}

func (auth *Auth) CheckPassword(user types.User, password string) bool {
	hashedPasswordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return hashedPasswordErr == nil
}

func (auth *Auth) EmailWhitelisted(emailSrc string) bool {
    if len(auth.OAuthWhitelist) == 0 {
        return true
    }
    for _, allowedPattern := range auth.OAuthWhitelist {
        // Check exact match first
        if allowedPattern == emailSrc {
            return true
        }
        // Check domain wildcard match
        if strings.HasPrefix(allowedPattern, "*@") {
            domain := strings.TrimPrefix(allowedPattern, "*@")
            if strings.HasSuffix(emailSrc, "@"+domain) {
                return true
            }
        }
    }
    return false
}

func (auth *Auth) CreateSessionCookie(c *gin.Context, data *types.SessionCookie) {
	log.Debug().Msg("Creating session cookie")
	sessions := sessions.Default(c)
	log.Debug().Msg("Setting session cookie")
	sessions.Set("username", data.Username)
	sessions.Set("provider", data.Provider)
	sessions.Save()
}

func (auth *Auth) DeleteSessionCookie(c *gin.Context) {
	log.Debug().Msg("Deleting session cookie")
	sessions := sessions.Default(c)
	sessions.Clear()
	sessions.Save()
}

func (auth *Auth) GetSessionCookie(c *gin.Context) (types.SessionCookie, error) {
	log.Debug().Msg("Getting session cookie")
	sessions := sessions.Default(c)

	cookieUsername := sessions.Get("username")
	cookieProvider := sessions.Get("provider")

	username, usernameOk := cookieUsername.(string)
	provider, providerOk := cookieProvider.(string)

	log.Debug().Str("username", username).Str("provider", provider).Msg("Parsed cookie")

	if !usernameOk || !providerOk {
		log.Warn().Msg("Session cookie invalid")
		return types.SessionCookie{}, nil
	}

	return types.SessionCookie{
		Username: username,
		Provider: provider,
	}, nil
}

func (auth *Auth) UserAuthConfigured() bool {
	return len(auth.Users) > 0
}
