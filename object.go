// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package uplink

import (
	"context"
	"time"

	"storj.io/common/storj"
)

// Object contains information about an object.
type Object struct {
	Key string

	Created time.Time
	Expires time.Time

	// TODO: Figure out what kind of size fields we want to have here
	// EncryptedSize int64

	User Metadata
}

// Metadata contains user metadata about the object.
type Metadata map[string]string

// Stat returns information about an object at the specific key.
func (project *Project) Stat(ctx context.Context, bucket, key string) (_ *Object, err error) {
	defer mon.Task()(&ctx)(&err)

	b := storj.Bucket{Name: bucket, PathCipher: storj.EncAESGCM}
	obj, err := project.db.GetObject(ctx, b, key)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	return convertObject(&obj), nil
}

// DeleteObject deletes the object at the specific key.
func (project *Project) DeleteObject(ctx context.Context, bucket, key string) (err error) {
	defer mon.Task()(&ctx)(&err)

	b := storj.Bucket{Name: bucket, PathCipher: storj.EncAESGCM}
	err = project.db.DeleteObject(ctx, b, key)

	return Error.Wrap(err)
}

// convertObject converts storj.Object to uplink.Object.
func convertObject(obj *storj.Object) *Object {
	return &Object{
		Key:     obj.Path,
		Created: obj.Created,
		Expires: obj.Expires,
		User:    obj.Metadata,
	}
}
