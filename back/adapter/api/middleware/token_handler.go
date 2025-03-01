package middleware

import (
	"errors"
	"net/http"
	"os"
)

const tokenName = "token_auth"

type CookieToken struct {
	token string
}

func NewCookieToken(token string) (*CookieToken, error) {
	if token == "" {
		return nil, errors.New("empty token")
	}

	return &CookieToken{token: token}, nil
}

func (c *CookieToken) Token() string {
	return c.token
}

func GetToken(req *http.Request) (*CookieToken, error) {
	var tokenString string
	token, err := req.Cookie(tokenName)
	if err != nil {
		tokenString = ""
	} else {
		tokenString = token.Value
	}

	tokenModel, err := NewCookieToken(tokenString)

	if err != nil {
		return nil, err
	}

	return tokenModel, nil
}

func SetToken(res http.ResponseWriter, token *CookieToken) {
	http.SetCookie(res, &http.Cookie{
		Name:     tokenName,
		Value:    token.Token(),
		HttpOnly: true,
		Secure:   os.Getenv("GO_ENVIRONMENT") == "production",
		MaxAge:   3600,
		Domain:   os.Getenv("DOMAIN_NAME"),
	})
}
