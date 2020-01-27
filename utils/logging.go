/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package utils

import (
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"
)

type logrusHook struct {
	t *testing.T
}

func (h logrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h logrusHook) Fire(entry *logrus.Entry) error {
	h.t.Helper()
	h.t.Log(entry.String())
	return nil
}

func NewTestingLogger(t *testing.T) *logrus.Logger {
	t.Helper()
	out := logrus.New()
	out.SetLevel(logrus.TraceLevel)
	out.SetOutput(ioutil.Discard)
	out.AddHook(logrusHook{t: t})
	return out
}
