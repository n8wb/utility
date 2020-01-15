/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

import "time"

type RunningInstance struct {
	Domain string `json:"domain"`
	Ports  []int  `json:"ports"`
}

type TestInfo struct {
	Name string `json:"name"`

	Instances []RunningInstance `json:"instances"`

	// Files contains file name -> url to fetch it
	Files       map[string]string `json:"files"`
	Environment map[string]string `json:"environment"`

	DefinitionID   string    `json:"definitionID" db:"definition_id"`
	OrganizationID string    `json:"organizationID" db:"organization_id"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	DestroyedAt    time.Time `json:"destroyedAt" db:"destroyed_at"`
	SpecFile       string    `json:"specFile"` //the definition spec
}
