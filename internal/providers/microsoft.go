package providers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"tinyauth/internal/constants"

	"github.com/rs/zerolog/log"
)

// Response for the Microsoft Graph user endpoint
type MicrosoftUserInfoResponse struct {
	Mail              string `json:"mail"`
	UserPrincipalName string `json:"userPrincipalName"`
	DisplayName       string `json:"displayName"`
	ID                string `json:"ID"`
}

// Response for the Microsoft Graph memberOf endpoint
type MicrosoftUserGroupsResponse struct {
	Value []MicrosoftGroup `json:"value"`
}

// Individual group object
type MicrosoftGroup struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

// The scopes required for the Microsoft provider
func MicrosoftScopes() []string {
	return []string{"openid", "profile", "email", "User.Read", "GroupMember.Read.All"}
}

func GetMicrosoftUser(client *http.Client, userURL ...string) (constants.Claims, error) {
	var user constants.Claims

	// Use custom user URL if provided, otherwise fall back to default
	url := "https://graph.microsoft.com/v1.0/me"
	if len(userURL) > 0 && userURL[0] != "" {
		url = userURL[0]
	}

	res, err := client.Get(url)
	if err != nil {
		return user, err
	}
	defer res.Body.Close()

	log.Debug().Msg("Got response from microsoft")

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return user, err
	}

	log.Debug().Msg("Read body from microsoft")

	var userInfo MicrosoftUserInfoResponse

	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return user, err
	}

	log.Debug().Msg("Attempt to parse user groups from microsoft")

	memberOfURL := fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%s/memberOf", userInfo.ID)
	groupRes, err := client.Get(memberOfURL)
	if err != nil {
		return user, err
	}
	defer groupRes.Body.Close()

	log.Debug().Msg("Got group response from microsoft")

	groupBody, err := io.ReadAll(groupRes.Body)
	if err != nil {
		return user, err
	}

	var groupResponse MicrosoftUserGroupsResponse
	if err := json.Unmarshal(groupBody, &groupResponse); err != nil {
		return user, err
	}

	// Collect group names
	groupNames := []any{}
	for _, g := range groupResponse.Value {
		groupNames = append(groupNames, g.DisplayName)
	}

	log.Debug().Msg("Parsed user from microsoft")

	// Prefer mail, fallback to UserPrincipalName
	email := userInfo.Mail
	if email == "" {
		email = userInfo.UserPrincipalName
	}
	user.PreferredUsername = strings.Split(email, "@")[0]
	user.Name = userInfo.DisplayName
	user.Email = email
	user.Groups = groupNames

	return user, nil
}
