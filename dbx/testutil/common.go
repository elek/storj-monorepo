// Copyright (C) 2019 Storj Labs, Inc.
// Copyright (C) 2017 Space Monkey, Inc.
// See LICENSE for copying information.

package testutil

import (
	"sort"
	"strings"
	"testing"
)

func Wrap(t *testing.T) *T {
	return &T{
		T: t,

		context: make(map[string]string),
	}
}

type T struct {
	*testing.T

	context map[string]string
}

func (t *T) Context(name string, val string) {
	t.context[name] = val
}

func (t *T) dumpContext() {
	var keys []string
	for key := range t.context {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		t.Logf("%s:\n%s", key, t.context[key])
	}
}

func (t *T) Assert(that bool) {
	t.Helper()
	if !that {
		t.dumpContext()
		t.Fatalf("expected condition to be true.")
	}
}

func (t *T) AssertNoError(err error) {
	t.Helper()
	if err != nil {
		t.dumpContext()
		t.Fatalf("expected no error. got %v", err)
	}
}

func (t *T) AssertError(err error, msg string) {
	t.Helper()
	if err == nil {
		t.dumpContext()
		t.Fatalf("expected an error containing %q. got nil", msg)
	}
	if !strings.Contains(err.Error(), msg) {
		t.dumpContext()
		t.Fatalf("expected an error containing %q. got %v", msg, err)
	}
}

func (t *T) AssertContains(msg, needle string) {
	t.Helper()
	if !strings.Contains(msg, needle) {
		t.dumpContext()
		t.Fatalf("expected the message to contain %q. got %s", needle, msg)
	}
}

func (t *T) Run(name string, f func(*T)) bool {
	return t.T.Run(name, func(t *testing.T) { f(Wrap(t)) })
}

func (t *T) Runp(name string, f func(*T)) bool {
	return t.T.Run(name, func(t *testing.T) { t.Parallel(); f(Wrap(t)) })
}
