// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fstesting

import (
	"os"
	"path"

	"github.com/jacobsa/gcloud/gcs/gcsutil"
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
)

////////////////////////////////////////////////////////////////////////
// Boilerplate
////////////////////////////////////////////////////////////////////////

type readOnlyTest struct {
	fsTest
}

func init() { registerSuitePrototype(&readOnlyTest{}) }

func (t *readOnlyTest) setUpFSTest(cfg FSTestConfig) {
	cfg.MountConfig.ReadOnly = true
	t.fsTest.setUpFSTest(cfg)
}

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *readOnlyTest) CreateObject() {
	AssertTrue(false, "TODO")
}

func (t *readOnlyTest) ModifyObject() {
	AssertTrue(false, "TODO")
}

func (t *readOnlyTest) DeleteObject() {
	// Create an object in the bucket.
	_, err := gcsutil.CreateObject(t.ctx, t.bucket, "foo", "taco")
	AssertEq(nil, err)

	// Attempt to delete it via the file system.
	err = os.Remove(path.Join(t.Dir, "foo"))
	ExpectThat(err, Error(HasSubstr("read-only")))

	// the bucket should not have been modified.
	contents, err := gcsutil.ReadObject(t.ctx, t.bucket, "foo")

	AssertEq(nil, err)
	ExpectEq("taco", contents)
}
