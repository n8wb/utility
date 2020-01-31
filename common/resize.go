/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

type Resize struct {
	Height uint   `json:"height"`
	Width  uint   `json:"width"`
	Host   string `json:"host"`
	ExecID string `json:"execID"`
}
