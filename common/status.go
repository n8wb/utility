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

type Status struct {
	Test      string `json:"test" db:"test_id"`
	Org       string `json:"org" db:"organization_id"`
	Def       string `json:"def" db:"definition_id"`
	Phase     string `json:"phase" db:"phase"`
	StepsLeft int    `json:"stepsLeft" db:"steps_left"`
	Message   string `json:"message,omitempty" db:"message"`
	Finished  bool   `json:"finished" db:"finished"`
}
