/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemconv(t *testing.T) {
	val, err := Memconv("20", 1000)
	assert.NoError(t, err)
	assert.Equal(t, int64(20000), val)

	val, err = Memconv("200 KB", 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(200*Kibi), val)

	val, err = Memconv("2000 KiB", 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(2000*Kibi), val)

	val, err = Memconv("2TB", 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(2*Tibi), val)

	val, err = Memconv("2GB", 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(2*Gibi), val)

	val, err = Memconv("2 MB", 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(2*Mibi), val)

	_, err = Memconv("foo", 0)
	assert.Error(t, err)

	val, err = Memconv("5mb", 0)
	assert.NoError(t, err)
	assert.Equal(t, int64(5*Mibi), val)
}
