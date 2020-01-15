/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

type Status struct {
	Test      string `json:"test" db:"test_id"`
	Org       string `json:"org" db:"organization_id"`
	Def       string `json:"def" db:"definition_id"`
	Phase     string `json:"phase" db:"phase"`
	StepsLeft int    `json:"stepsLeft" db:"steps_left"`
	Message   string `json:"message,omitempty" db:"message"`
	Finished  bool   `json:"finished" db:"finished"`
}
