package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sort"

	"github.com/Pratchaya0/auth-api-gin/usecases"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/line"
	"github.com/markbates/goth/providers/tiktok"
	"github.com/markbates/goth/providers/twitter"
	"github.com/markbates/goth/providers/twitterv2"
)

var indexTemplate = `{{range $key,$value:=.Providers}}
    <p><a href="/auth/{{$value}}">Log in with {{index $.ProvidersMap $value}}</a></p>
{{end}}`

var userTemplate = `
<p><a href="/auth/logout/{{.Provider}}">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

var providerIndex *ProviderIndex

func Provider() *ProviderIndex {
	return providerIndex
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	goth.UseProviders(
		// Use twitterv2 instead of twitter if you only have access to the Essential API Level
		// the twitter provider uses a v1.1 API that is not available to the Essential Level
		twitterv2.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitterv2/callback"),
		// If you'd like to use authenticate instead of authorize in TwitterV2 provider, use this instead.
		// twitterv2.NewAuthenticate(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitterv2/callback"),

		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),
		// If you'd like to use authenticate instead of authorize in Twitter provider, use this instead.
		// twitter.NewAuthenticate(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),

		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:8080/auth/github/callback"),
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:3000/auth/google/callback"),
		tiktok.New(os.Getenv("TIKTOK_KEY"), os.Getenv("TIKTOK_SECRET"), "http://localhost:3000/auth/tiktok/callback"),
		line.New(os.Getenv("LINE_KEY"), os.Getenv("LINE_SECRET"), "http://localhost:3000/auth/line/callback", "profile", "openid", "email"),
	)
	gothic.Store = cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))

	m := map[string]string{
		"google":    "Google",
		"github":    "Github",
		"tiktok":    "TikTok",
		"line":      "LINE",
		"twitter":   "Twitter",
		"twitterv2": "Twitter",
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	providerIndex = &ProviderIndex{Providers: keys, ProvidersMap: m}
}

type Auth2Controller struct {
	auth2UseCase *usecases.UserUseCase
}

func NewAuth2Controller(useCase *usecases.UserUseCase) *Auth2Controller {
	return &Auth2Controller{auth2UseCase: useCase}
}

func (ctrl *Auth2Controller) OAuthIndex(c *gin.Context) {
	t, _ := template.New("foo").Parse(indexTemplate)
	t.Execute(c.Writer, providerIndex)
}

func (ctrl *Auth2Controller) OAuthStart(c *gin.Context) {
	req := c.Request.URL.Query()
	req.Add("provider", c.Param("provider"))
	c.Request.URL.RawQuery = req.Encode()

	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(c.Writer, gothUser)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

func (ctrl *Auth2Controller) OAuthLogout(c *gin.Context) {
	req := c.Request.URL.Query()
	req.Add("provider", c.Param("provider"))
	c.Request.URL.RawQuery = req.Encode()

	gothic.Logout(c.Writer, c.Request)
	c.Writer.Header().Set("Location", "/")
	c.Writer.WriteHeader(http.StatusTemporaryRedirect)
}

func (ctrl *Auth2Controller) OAuthCallback(c *gin.Context) {
	req := c.Request.URL.Query()
	req.Add("provider", c.Param("provider"))
	c.Request.URL.RawQuery = req.Encode()

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Fprintln(c.Writer, err)
		return
	}
	t, _ := template.New("foo").Parse(userTemplate)
	t.Execute(c.Writer, user)
}
