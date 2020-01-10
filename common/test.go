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

// Test represents a whiteblock test
type Test struct {
	ID             string    `json:"id" db:"id"`
	DefinitionID   string    `json:"definitionID" db:"definition_id"`
	OrganizationID string    `json:"organizationID" db:"organization_id"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	DestroyedAt    time.Time `json:"destroyedAt" db:"destroyed_at"`
}