/*
	Copyright 2019 whiteblock Inc.
	This file is a part of the utility.

	Utility is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	Utility is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

var (
	ErrMissingJWT = errors.New("missing JWT in authorization header")
	InvalidHeader = errors.New("invalid auth header")
	InvalidClaim  = errors.New("invalid claim")
)

type UserContext struct {
	KeycloakID string `json:"sub"`
	Username   string `json:"preferred_username"`
	Email      string `json:"email"`
}

func GetUserContext(token *oidc.IDToken) (*UserContext, error) {
	userContext := &UserContext{}

	err := token.Claims(userContext)
	if err != nil {
		log.Errorf("%s: %s", InvalidClaim, err)
		return nil, InvalidClaim
	}
	return userContext, nil
}

func VerifyToken(verifier *oidc.IDTokenVerifier, ctx context.Context, header string) (*oidc.IDToken, error) {
	bearerToken := strings.Split(header, "Bearer ")
	tokenString := bearerToken[1]

	token, err := verifier.Verify(ctx, tokenString)
	if err != nil {
		log.Errorf("Failed to verify ID Token: %s", err)
	}
	return token, nil
}

// TODO deprecated
// ExtractJwt will attempt to extract and return the jwt from the auth header
func ExtractJwt(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")

	if len(tokenString) == 0 {
		return "", ErrMissingJWT
	}
	splt := strings.Split(tokenString, " ")
	if len(splt) < 2 {
		return "", InvalidHeader
	}
	return splt[1], nil
}

// TODO deprecated
func GetValidatedToken(publicKey string, access_token string) (*jwt.Token, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TODO deprecated
func GetUserInfo(token *jwt.Token) (*UserContext, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userInfo := &UserContext{
			KeycloakID: claims["sub"].(string),
			Username:   claims["preferred_username"].(string),
			Email:      claims["email"].(string),
		}
		return userInfo, nil
	} else {
		return nil, nil
	}
}