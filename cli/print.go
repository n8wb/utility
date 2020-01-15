/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
package cli

import (
	"encoding/json"
	"fmt"
	"github.com/Whiteblock/go-prettyjson"
	"os"
)

func prettypi(i interface{}) string {
	_, noPretty := os.LookupEnv("NO_PRETTY")
	if noPretty {
		out, _ := json.Marshal(i)
		return string(out)
	}
	out, _ := prettyjson.Marshal(i)
	return string(out)
}

func Print(i interface{}) {
	switch i.(type) {
	case string:
		_, noPretty := os.LookupEnv("NO_PRETTY")
		if noPretty {
			fmt.Println(i.(string))
		} else {
			fmt.Printf("\033[97m%s\033[0m\n", i.(string))
		}
	default:
		fmt.Println(prettypi(i))
	}
}

func Printf(format string, a ...interface{}) {
	Print(fmt.Sprintf(format, a...))
}
