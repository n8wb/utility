/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

import "time"

type TestInfo struct {
	Name string `json:"name"`

	Domains []string `json:"domains"` //domain names

	// Ports contains the exposed ports,
	// len(Ports) == len(Domains), and Domains[n] should have all of the
	// ports in Ports[n] exposed. (Note: may not be always up, but will be up atleast once)
	Ports [][]int `json:"ports"` //exposed ports,

	// Files contains file name -> url to fetch it
	Files       map[string]string `json:"files"`
	Environment map[string]string `json:"environment"`

	DefinitionID   string    `json:"definitionID" db:"definition_id"`
	OrganizationID string    `json:"organizationID" db:"organization_id"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	DestroyedAt    time.Time `json:"destroyedAt" db:"destroyed_at"`
	SpecFile       []byte    `json:"specFile"` //the definition spec
}
