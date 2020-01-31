/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

type Exec struct {
	Test        string   `json:"test"`       //testid
	Target      string   `json:"target"`     //the entity to attach to
	Command     []string `json:"command"`    //the command to run
	Privileged  bool     `json:"privileged"` //should extra privileges be granted
	Interactive bool     `json:"interactive"`
	TTY         bool     `json:"tty"`
	Detach      bool     `json:"detach"`
}

type ExecInfo struct {
	Test string `json:"test"`
	ID   string `json:"id"`
	Host string `json:"host"`
}

type ExecAttach struct {
	ExecInfo
	TTY    bool `json:"tty"`
	Detach bool `json:"detach"`
}
