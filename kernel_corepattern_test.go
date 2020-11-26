// Copyright 2019 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !windows

package procfs

import (
	"bytes"
	"testing"
)

func TestKernelCorePattern(t *testing.T) {
	have, err := getProcFixtures(t).KernelCorePattern()
	if err != nil {
		t.Fatal(err)
	}

	want := []byte("|/usr/share/apport/apport %p %s %c %d %P %E")
	if !bytes.Equal(want, have) {
		t.Fatalf("want pattern `%s`, have `%s`", want, have)
	}

}
