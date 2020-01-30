/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

import "time"

type User struct {
	ID         string    `json:"id" db:"id"`
	KeycloakID string    `json:"keycloak_id" db:"keycloak_id"`
	Username   string    `json:"username" db:"username"`
	Email      string    `json:"email" db:"email"`
	CreatedAt  time.Time `json:"created_at,omitonempty" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitonempty" db:"updated_at"`
}
