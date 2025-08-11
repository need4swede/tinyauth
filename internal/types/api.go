package types

// LoginQuery is the query parameters for the login endpoint
type LoginQuery struct {
	RedirectURI string `url:"redirect_uri"`
}

// LoginRequest is the request body for the login endpoint
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// OAuthRequest is the request for the OAuth endpoint
type OAuthRequest struct {
	Provider string `uri:"provider" binding:"required"`
}

// UnauthorizedQuery is the query parameters for the unauthorized endpoint
type UnauthorizedQuery struct {
	Username string `url:"username"`
	Resource string `url:"resource"`
	GroupErr bool   `url:"groupErr"`
	IP       string `url:"ip"`
}

// Proxy is the uri parameters for the proxy endpoint
type Proxy struct {
	Proxy string `uri:"proxy" binding:"required"`
}

// User Context response is the response for the user context endpoint
type UserContextResponse struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	IsLoggedIn  bool   `json:"isLoggedIn"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Provider    string `json:"provider"`
	Oauth       bool   `json:"oauth"`
	TotpPending bool   `json:"totpPending"`
}

// App Context is the response for the app context endpoint
type AppContext struct {
	Status                int      `json:"status"`
	Message               string   `json:"message"`
	ConfiguredProviders   []string `json:"configuredProviders"`
	DisableContinue       bool     `json:"disableContinue"`
	Title                 string   `json:"title"`
	GenericName           string   `json:"genericName"`
	Domain                string   `json:"domain"`
	ForgotPasswordMessage string   `json:"forgotPasswordMessage"`
	BackgroundImage       string   `json:"backgroundImage"`
	LoginTitle            string   `json:"loginTitle"`
	LoginSubtitle         string   `json:"loginSubtitle"`
	UsernameTitle         string   `json:"usernameTitle"`
	PasswordTitle         string   `json:"passwordTitle"`
	UsernamePlaceholder   string   `json:"usernamePlaceholder"`
	PasswordPlaceholder   string   `json:"passwordPlaceholder"`
	Logo                  string   `json:"logo"`
	LogoSize              string   `json:"logoSize"`
	LoginTitleSize        string   `json:"loginTitleSize"`
	LoginSubtitleSize     string   `json:"loginSubtitleSize"`
	DisableBorder         bool     `json:"disableBorder"`
	LoginCardColor        string   `json:"loginCardColor"`
	LoginTitleColor       string   `json:"loginTitleColor"`
	LoginSubtitleColor    string   `json:"loginSubtitleColor"`
	UsernameTextColor     string   `json:"usernameTextColor"`
	PasswordTextColor     string   `json:"passwordTextColor"`
	GoogleButtonText           string `json:"googleButtonText"`
	GoogleButtonTextColor      string `json:"googleButtonTextColor"`
	GoogleButtonBackgroundColor string `json:"googleButtonBackgroundColor"`
	GoogleButtonHideIcon       bool   `json:"googleButtonHideIcon"`
	MicrosoftButtonText           string `json:"microsoftButtonText"`
	MicrosoftButtonTextColor      string `json:"microsoftButtonTextColor"`
	MicrosoftButtonBackgroundColor string `json:"microsoftButtonBackgroundColor"`
	MicrosoftButtonHideIcon       bool   `json:"microsoftButtonHideIcon"`
	GithubButtonText           string `json:"githubButtonText"`
	GithubButtonTextColor      string `json:"githubButtonTextColor"`
	GithubButtonBackgroundColor string `json:"githubButtonBackgroundColor"`
	GithubButtonHideIcon       bool   `json:"githubButtonHideIcon"`
	GenericButtonText           string `json:"genericButtonText"`
	GenericButtonTextColor      string `json:"genericButtonTextColor"`
	GenericButtonBackgroundColor string `json:"genericButtonBackgroundColor"`
	GenericButtonHideIcon       bool   `json:"genericButtonHideIcon"`
	LoginButtonText           string `json:"loginButtonText"`
	LoginButtonTextColor      string `json:"loginButtonTextColor"`
	LoginButtonBackgroundColor string `json:"loginButtonBackgroundColor"`
	OAuthAutoRedirect     string   `json:"oauthAutoRedirect"`
}

// Totp request is the request for the totp endpoint
type TotpRequest struct {
	Code string `json:"code"`
}
