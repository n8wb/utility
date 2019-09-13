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
