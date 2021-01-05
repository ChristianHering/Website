package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"net/http"
	"net/url"

	"github.com/coreos/go-oidc"
	"github.com/pkg/errors"

	"golang.org/x/oauth2"

	"github.com/gorilla/sessions"
)

//SessionStore Session cookie store
var SessionStore *sessions.FilesystemStore

//Authenticator Auth0 default authenticator struct
type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

//Creates a filesystem session store
//for user request authentication
func setupAuth() {
	SessionStore = sessions.NewFilesystemStore("", Config.AuthenticationConfig.CookieStoreKeys...)
	gob.Register(map[string]interface{}{})
	return
}

//NewAuthenticator Instantiates a new OpenID/OAuth
//client, then returns an authenticator struct
func NewAuthenticator(requestHost string) (*Authenticator, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://"+Config.AuthenticationConfig.Auth0Domain+"/")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	conf := oauth2.Config{
		ClientID:     Config.AuthenticationConfig.Auth0ClientID,
		ClientSecret: Config.AuthenticationConfig.Auth0ClientSecret,
		RedirectURL:  "https://" + requestHost + "/callback",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}

//LoginHandler Sets up a new OAuth/OpenID client, then redirects our user
//to Auth0 for authentication. User sent back via CallbackHandler()
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 32)
	_, err := rand.Read(b) //Generate random state
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}
	state := base64.StdEncoding.EncodeToString(b)

	session, err := SessionStore.Get(r, "auth-session")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}

	authenticator, err := NewAuthenticator(r.Host)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}

	http.Redirect(w, r, authenticator.Config.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

//LogoutHandler Removes a user's login token, and clears
//their session stored locally to fully log the user out
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	logoutURL, err := url.Parse("https://" + Config.AuthenticationConfig.Auth0Domain)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}

	logoutURL.Path += "/v2/logout"
	parameters := url.Values{}

	returnTo, err := url.Parse("https://" + r.Host)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", Config.AuthenticationConfig.Auth0ClientID)
	logoutURL.RawQuery = parameters.Encode()

	http.Redirect(w, r, logoutURL.String(), http.StatusTemporaryRedirect)
}

//CallbackHandler This is where users are directed after authentication on Auth0
//
//Call this with standard middleware, then use our
//authentication handler to protect other handlers
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	session, err := SessionStore.Get(r, "auth-session")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}

	if r.URL.Query().Get("state") != session.Values["state"] {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authenticator, err := NewAuthenticator(r.Host)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}

	token, err := authenticator.Config.Exchange(context.TODO(), r.URL.Query().Get("code"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		panic(errors.WithStack(err))
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}

	oidcConfig := &oidc.Config{
		ClientID: Config.AuthenticationConfig.Auth0ClientID,
	}

	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(context.TODO(), rawIDToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}

	var profile map[string]interface{} //Getting the userInfo3
	if err := idToken.Claims(&profile); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}

	session.Values["id_token"] = rawIDToken
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile
	err = session.Save(r, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(errors.WithStack(err))
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther) //User's now authenticated. Send to user page.
}
