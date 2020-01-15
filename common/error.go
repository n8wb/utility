/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
package common

import (
	"fmt"
	"runtime/debug"

	"github.com/jinzhu/copier"
)

// ErrorType represents the types of errors
type ErrorType int

const (
	// Default is the generic error type
	Default ErrorType = iota

	// Fatal is the type for errors which are not recoverable
	Fatal
)

// Error is an improve error type for the inclusion of more information
type Error struct {
	Type    ErrorType              `json:"type"`
	Message string                 `json:"message"`
	Meta    map[string]interface{} `json:"meta"`
}

// NewError creates a new error of the default type
func NewError(msg interface{}) Error {
	return Error{
		Type:    Default,
		Message: fmt.Sprint(msg),
		Meta:    map[string]interface{}{},
	}
}

// NewFatalError creates an error that does not warrant a retry
func NewFatalError(msg interface{}) Error {
	out := NewError(msg)
	out.Type = Fatal
	return out
}

// Error gives the error message as a string
func (err Error) Error() string {
	return err.Message
}

// InjectMeta allows the insertion of meta into the error
func (err Error) InjectMeta(meta map[string]interface{}) Error {
	if err.Meta == nil {
		copier.Copy(&err.Meta, meta)
		return err
	}
	for k, v := range meta {
		err.Meta[k] = v
	}
	return err
}

func (err Error) WithStack() Error {
	return err.InjectMeta(map[string]interface{}{
		"stack": debug.Stack(),
	})
}

func (err Error) MarshalText() (text []byte, _ error) {
	text = []byte(fmt.Sprintf("%+v", err))
	return
}

// IsFatalError checks whether or not the error is a fatal error
func IsFatalError(err error) bool {
	if err == nil {
		return false
	}
	err2, isErr := err.(*Error)
	if !isErr {
		return false
	}

	return err2.Type == Fatal
}
