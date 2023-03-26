// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"storj.io/picobuf/internal/sizebench/basz"
	"storj.io/picobuf/internal/sizebench/pico"
	"storj.io/picobuf/internal/sizebench/pico/sml"
)

func main() {
	fmt.Println([]byte{})
	basz.CallDynamicMethod(fmt.Println)
	pico.Test(&sml.Nop{})
	pico.Test(&sml.Types{})
	basz.CallDynamicMethod(&sml.Types{})
}
