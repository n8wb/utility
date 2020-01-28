/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

import (
	"database/sql"
)

// Instance represents a virtual machine in the cloud
type Instance struct {
	Provider  string         `json:"provider" db:"provider"`
	Project   string         `json:"project" db:"project"`
	Zone      string         `json:"zone" db:"zone"`
	Name      string         `json:"name" db:"name"`
	BiomeID   string         `json:"biomeId" db:"biome_id"`
	IP        string         `json:"ip" db:"ip"`
	HumanName sql.NullString `json:"humanName,omitempty" db:"human_name"`
}
