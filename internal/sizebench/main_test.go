// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"testing"
)

func TestMessageSize(t *testing.T) {
	arch := OsArch{Os: "linux", Arch: "amd64"}

	pkg, err := analyzePkg(t.TempDir(), arch, "pico/one")
	if err != nil {
		t.Fatal(err)
	}

	checkSize(t, "message size", pkg.size, 67818, 4000)
	checkSize(t, "encode size", pkg.symbols.sum("one.(*Types).Encode"), 1306, 64)
	checkSize(t, "decode size", pkg.symbols.sum("one.(*Types).Decode"), 1834, 64)
}

func checkSize(t *testing.T, name string, size, target, eps int64) {
	t.Helper()

	low := target - eps
	high := target + eps

	if size < low || high < size {
		t.Errorf("%s %v not in range %v±1024 (low %v, high %v)", name,
			memorySize(size), memorySize(target),
			memorySize(low), memorySize(high),
		)
	} else {
		t.Logf("%s %v", name, memorySize(size))
	}
}
