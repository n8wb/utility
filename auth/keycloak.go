/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
package auth

import (
	"context"
	"errors"
	"fmt"
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
	if token == nil {
		return nil, ErrMissingJWT
	}
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
		return nil, fmt.Errorf("Failed to verify ID Token: %s", err)
	}
	return token, nil
}
