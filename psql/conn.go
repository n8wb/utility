/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package psql

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func OpenConnectionPool(conf Config, log logrus.Ext1FieldLogger) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName)
	log.Info("Trying to connect to database...")
	connection, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Couldn't connect to database", err)
	}
	connection.SetMaxOpenConns(conf.MaxOpenConns)
	connection.SetMaxIdleConns(conf.MaxIdleConns)
	connection.SetConnMaxLifetime(time.Duration(conf.MaxLifetimeConns) * time.Second)

	log.Info("Database connection is successful")
	return connection, nil
}
