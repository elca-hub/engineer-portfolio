package token_auth

import (
	"net/http"
	"os"
)

const tokenName = "token_auth"

func GetToken(req *http.Request) string {
	var tokenString string
	token, err := req.Cookie(tokenName)
	if err != nil {
		tokenString = ""
	} else {
		tokenString = token.Value
	}

	return tokenString
}

func SetToken(res http.ResponseWriter, token string) {
	http.SetCookie(res, &http.Cookie{
		Name:     tokenName,
		Value:    token,
		HttpOnly: true,
		Secure:   os.Getenv("GO_ENVIRONMENT") == "production",
		MaxAge:   3600,
		Domain:   os.Getenv("DOMAIN_NAME"),
	})
}
