// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"

	"storj.io/picobuf/internal/sizebench/prot"
	"storj.io/picobuf/internal/sizebench/prot/sml"
)

func main() {
	fmt.Println([]byte{})
	prot.Test(&sml.Nop{})
	prot.Test(&sml.Types{})
}
