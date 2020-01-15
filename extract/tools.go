/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package extract

import (
	"firebase.google.com/go/auth"
	"fmt"
	"net/http"
)

func SafeTokenExtract(r *http.Request) (*auth.Token, error) {
	if r.Context().Value("token") == nil {
		return nil, fmt.Errorf("missing firebase auth token")
	}
	token, ok := r.Context().Value("token").(*auth.Token)
	if !ok {
		return nil, fmt.Errorf("token is of incorrect type")
	}
	return token, nil
}

func SafeEmailExtract(token *auth.Token) (string, error) {
	_, exists := token.Claims["email"]
	if !exists {
		return "", fmt.Errorf("missing email in claims")
	}
	email, typeIsRight := token.Claims["email"].(string)
	if !typeIsRight {
		return "", fmt.Errorf("misformed type for email, expected string")
	}
	return email, nil
}
