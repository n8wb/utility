/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
package utils

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/Pallinder/go-randomdata"
	log "github.com/sirupsen/logrus"
	"github.com/whiteblock/go.uuid"
)

const (
	_          = iota
	Kibi int64 = 1 << (10 * iota)
	Mibi
	Gibi
	Tibi
)

//GetUUIDString generates a new UUID
func GetUUIDString() string {
	uid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return uid.String()
}

func LogErrorN(err error, n int) error {
	if err == nil {
		return err // do nothing if the given err is nil
	}
	_, file, line, ok := runtime.Caller(n)
	if !ok {
		log.Error(err.Error())
	} else {
		log.WithFields(log.Fields{"file": file, "line": line}).Error(err.Error())
	}

	return err
}

// LogError takes in an error, logs that error and returns that error.
// Used to help reduce code clutter and unify the error handling in the code.
// Has no effect if err == nil
func LogError(err error) error {
	return LogErrorN(err, 2)
}

func Memconv(mem string, defaultMulti int64) (int64, error) {

	m := strings.ToLower(mem)
	m = strings.Replace(m, " ", "", -1)
	m = strings.Replace(m, "ib", "b", -1)

	var multiplier int64 = defaultMulti

	if strings.HasSuffix(m, "kb") || strings.HasSuffix(m, "k") {
		multiplier = Kibi
	} else if strings.HasSuffix(m, "mb") || strings.HasSuffix(m, "m") {
		multiplier = Mibi
	} else if strings.HasSuffix(m, "gb") || strings.HasSuffix(m, "g") {
		multiplier = Gibi
	} else if strings.HasSuffix(m, "tb") || strings.HasSuffix(m, "t") {
		multiplier = Tibi
	}

	i, err := strconv.ParseInt(strings.Trim(m, "bgkmti"), 10, 64)
	if err != nil {
		return -1, err
	}

	return i * multiplier, nil
}

func CreateDNSSubdomains(n int) (dns []string) {
	for i := 0; i < n; i++ {
		dns = append(dns, strings.ToLower(randomdata.SillyName()))
	}
	return
}
