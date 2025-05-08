package auth

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"
	"tinyauth/internal/docker"
	"tinyauth/internal/types"
	"tinyauth/internal/utils"

	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func NewAuth(config types.AuthConfig, docker *docker.Docker) *Auth {
	return &Auth{
		Config:        config,
		Docker:        docker,
		LoginAttempts: make(map[string]*types.LoginAttempt),
	}
}

type Auth struct {
	Config        types.AuthConfig
	Docker        *docker.Docker
	LoginAttempts map[string]*types.LoginAttempt
	LoginMutex    sync.RWMutex
}

func (auth *Auth) GetSession(c *gin.Context) (*sessions.Session, error) {
	// Create cookie store
	store := sessions.NewCookieStore([]byte(auth.Config.Secret))

	// Configure cookie store
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   auth.Config.SessionExpiry,
		Secure:   auth.Config.CookieSecure,
		HttpOnly: true,
		Domain:   fmt.Sprintf(".%s", auth.Config.Domain),
	}

	// Get session
	session, err := store.Get(c.Request, "tinyauth")
	if err != nil {
		log.Error().Err(err).Msg("Failed to get session")
		return nil, err
	}

	return session, nil
}

func (auth *Auth) GetUser(username string) *types.User {
	// Loop through users and return the user if the username matches
	for _, user := range auth.Config.Users {
		if user.Username == username {
			return &user
		}
	}
	return nil
}

func (auth *Auth) CheckPassword(user types.User, password string) bool {
	// Compare the hashed password with the password provided
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}

// IsAccountLocked checks if a username or IP is locked due to too many failed login attempts
func (auth *Auth) IsAccountLocked(identifier string) (bool, int) {
	auth.LoginMutex.RLock()
	defer auth.LoginMutex.RUnlock()

	// Return false if rate limiting is not configured
	if auth.Config.LoginMaxRetries <= 0 || auth.Config.LoginTimeout <= 0 {
		return false, 0
	}

	// Check if the identifier exists in the map
	attempt, exists := auth.LoginAttempts[identifier]
	if !exists {
		return false, 0
	}

	// If account is locked, check if lock time has expired
	if attempt.LockedUntil.After(time.Now()) {
		// Calculate remaining lockout time in seconds
		remaining := int(time.Until(attempt.LockedUntil).Seconds())
		return true, remaining
	}

	// Lock has expired
	return false, 0
}

// RecordLoginAttempt records a login attempt for rate limiting
func (auth *Auth) RecordLoginAttempt(identifier string, success bool) {
	// Skip if rate limiting is not configured
	if auth.Config.LoginMaxRetries <= 0 || auth.Config.LoginTimeout <= 0 {
		return
	}

	auth.LoginMutex.Lock()
	defer auth.LoginMutex.Unlock()

	// Get current attempt record or create a new one
	attempt, exists := auth.LoginAttempts[identifier]
	if !exists {
		attempt = &types.LoginAttempt{}
		auth.LoginAttempts[identifier] = attempt
	}

	// Update last attempt time
	attempt.LastAttempt = time.Now()

	// If successful login, reset failed attempts
	if success {
		attempt.FailedAttempts = 0
		attempt.LockedUntil = time.Time{} // Reset lock time
		return
	}

	// Increment failed attempts
	attempt.FailedAttempts++

	// If max retries reached, lock the account
	if attempt.FailedAttempts >= auth.Config.LoginMaxRetries {
		attempt.LockedUntil = time.Now().Add(time.Duration(auth.Config.LoginTimeout) * time.Second)
		log.Warn().Str("identifier", identifier).Int("timeout", auth.Config.LoginTimeout).Msg("Account locked due to too many failed login attempts")
	}
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

func (auth *Auth) CreateSessionCookie(c *gin.Context, data *types.SessionCookie) error {
	log.Debug().Msg("Creating session cookie")

	// Get session
	session, err := auth.GetSession(c)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get session")
		return err
	}

	log.Debug().Msg("Setting session cookie")

	// Calculate expiry
	var sessionExpiry int

	if data.TotpPending {
		sessionExpiry = 3600
	} else {
		sessionExpiry = auth.Config.SessionExpiry
	}

	// Set data
	session.Values["username"] = data.Username
	session.Values["name"] = data.Name
	session.Values["email"] = data.Email
	session.Values["provider"] = data.Provider
	session.Values["expiry"] = time.Now().Add(time.Duration(sessionExpiry) * time.Second).Unix()
	session.Values["totpPending"] = data.TotpPending
	session.Values["oauthGroups"] = data.OAuthGroups

	// Save session
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		log.Error().Err(err).Msg("Failed to save session")
		return err
	}

	// Return nil
	return nil
}

func (auth *Auth) DeleteSessionCookie(c *gin.Context) error {
	log.Debug().Msg("Deleting session cookie")

	// Get session
	session, err := auth.GetSession(c)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get session")
		return err
	}

	// Delete all values in the session
	for key := range session.Values {
		delete(session.Values, key)
	}

	// Save session
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		log.Error().Err(err).Msg("Failed to save session")
		return err
	}

	// Return nil
	return nil
}

func (auth *Auth) GetSessionCookie(c *gin.Context) (types.SessionCookie, error) {
	log.Debug().Msg("Getting session cookie")

	// Get session
	session, err := auth.GetSession(c)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get session")
		return types.SessionCookie{}, err
	}

	log.Debug().Msg("Got session")

	// Get data from session
	username, usernameOk := session.Values["username"].(string)
	email, emailOk := session.Values["email"].(string)
	name, nameOk := session.Values["name"].(string)
	provider, providerOK := session.Values["provider"].(string)
	expiry, expiryOk := session.Values["expiry"].(int64)
	totpPending, totpPendingOk := session.Values["totpPending"].(bool)
	oauthGroups, oauthGroupsOk := session.Values["oauthGroups"].(string)

	if !usernameOk || !providerOK || !expiryOk || !totpPendingOk || !emailOk || !nameOk || !oauthGroupsOk {
		log.Warn().Msg("Session cookie is invalid")

		// If any data is missing, delete the session cookie
		auth.DeleteSessionCookie(c)

		// Return empty cookie
		return types.SessionCookie{}, nil
	}

	// Check if the cookie has expired
	if time.Now().Unix() > expiry {
		log.Warn().Msg("Session cookie expired")

		// If it has, delete it
		auth.DeleteSessionCookie(c)

		// Return empty cookie
		return types.SessionCookie{}, nil
	}

	log.Debug().Str("username", username).Str("provider", provider).Int64("expiry", expiry).Bool("totpPending", totpPending).Str("name", name).Str("email", email).Str("oauthGroups", oauthGroups).Msg("Parsed cookie")

	// Return the cookie
	return types.SessionCookie{
		Username:    username,
		Name:        name,
		Email:       email,
		Provider:    provider,
		TotpPending: totpPending,
		OAuthGroups: oauthGroups,
	}, nil
}

func (auth *Auth) UserAuthConfigured() bool {
	// If there are users, return true
	return len(auth.Config.Users) > 0
}

func (auth *Auth) ResourceAllowed(c *gin.Context, context types.UserContext, labels types.TinyauthLabels) bool {
	// Check if oauth is allowed
	if context.OAuth {
		log.Debug().Msg("Checking OAuth whitelist")
		return utils.CheckWhitelist(labels.OAuthWhitelist, context.Email)
	}

	// Check users
	log.Debug().Msg("Checking users")

	return utils.CheckWhitelist(labels.Users, context.Username)
}

func (auth *Auth) OAuthGroup(c *gin.Context, context types.UserContext, labels types.TinyauthLabels) bool {
	// Check if groups are required
	if labels.OAuthGroups == "" {
		return true
	}

	// Check if we are using the generic oauth provider
	if context.Provider != "generic" {
		log.Debug().Msg("Not using generic provider, skipping group check")
		return true
	}

	// Split the groups by comma (no need to parse since they are from the API response)
	oauthGroups := strings.Split(context.OAuthGroups, ",")

	// For every group check if it is in the required groups
	for _, group := range oauthGroups {
		if utils.CheckWhitelist(labels.OAuthGroups, group) {
			log.Debug().Str("group", group).Msg("Group is in required groups")
			return true
		}
	}

	// No groups matched
	log.Debug().Msg("No groups matched")

	// Return false
	return false
}

func (auth *Auth) AuthEnabled(c *gin.Context, labels types.TinyauthLabels) (bool, error) {
	// Get headers
	uri := c.Request.Header.Get("X-Forwarded-Uri")

	// Check if the allowed label is empty
	if labels.Allowed == "" {
		// Auth enabled
		return true, nil
	}

	// Compile regex
	regex, err := regexp.Compile(labels.Allowed)

	// If there is an error, invalid regex, auth enabled
	if err != nil {
		log.Warn().Err(err).Msg("Invalid regex")
		return true, err
	}

	// Check if the uri matches the regex
	if regex.MatchString(uri) {
		// Auth disabled
		return false, nil
	}

	// Auth enabled
	return true, nil
}

func (auth *Auth) GetBasicAuth(c *gin.Context) *types.User {
	// Get the Authorization header
	username, password, ok := c.Request.BasicAuth()

	// If not ok, return an empty user
	if !ok {
		return nil
	}

	// Return the user
	return &types.User{
		Username: username,
		Password: password,
	}
}
