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
	"strings"

	"github.com/coreos/go-oidc"
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
