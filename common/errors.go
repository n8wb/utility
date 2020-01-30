/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

import "errors"

var EmptyDBResult = errors.New("not found")
var AlreadyExists = errors.New("already exists")
var AccessDenied = errors.New("access denied")
var NotAMember = errors.New("not a member of this org")
var ValidationError = errors.New("validation error")
