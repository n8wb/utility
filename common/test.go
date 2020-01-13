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

package common

import (
	"database/sql"
	"encoding/json"
	"time"
)

type TestDefinition struct {
	ID         string    `json:"id" db:"id"`
	KeycloakID string    `json:"userID" db:"keycloak_id"`
	OrgID      string    `json:"orgID" db:"organization_id"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
}

// Test represents a whiteblock test
type Test struct {
	ID             string    `json:"id" db:"id"`
	DefinitionID   string    `json:"definitionID" db:"definition_id"`
	OrganizationID string    `json:"organizationID" db:"organization_id"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	Status         string    `json:"status" db:"status"`
	DestroyedAt    time.Time `json:"destroyedAt" db:"destroyed_at"`
	Name           string    `json:"name" db:"name"`
}

// Test represents a whiteblock test
type TestSQL struct {
	ID             string         `json:"id" db:"id"`
	DefinitionID   string         `json:"definitionID" db:"definition_id"`
	OrganizationID string         `json:"organizationID" db:"organization_id"`
	CreatedAt      sql.NullTime   `json:"createdAt" db:"created_at"`
	Status         sql.NullString `json:"status" db:"status"`
	DestroyedAt    sql.NullTime   `json:"destroyedAt" db:"destroyed_at"`
	Name           sql.NullString `json:"name" db:"name"`
}

func (ts TestSQL) Test() Test {
	out := Test{
		ID:             ts.ID,
		DefinitionID:   ts.DefinitionID,
		OrganizationID: ts.OrganizationID,
	}
	if ts.CreatedAt.Valid {
		out.CreatedAt = ts.CreatedAt.Time
	}

	if ts.Status.Valid {
		out.Status = ts.Status.String
	}

	if ts.DestroyedAt.Valid {
		out.DestroyedAt = ts.DestroyedAt.Time
	}

	if ts.Name.Valid {
		out.Name = ts.Name.String
	}
	return out
}

func (ts TestSQL) MarshalJSON() ([]byte, error) {
	return json.Marshal(ts.Test())
}
