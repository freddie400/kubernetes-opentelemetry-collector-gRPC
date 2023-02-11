// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pcommon

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
)

func TestTraceState_MoveTo(t *testing.T) {
	ms := TraceState(internal.GenerateTestTraceState())
	dest := NewTraceState()
	ms.MoveTo(dest)
	assert.Equal(t, NewTraceState(), ms)
	assert.Equal(t, TraceState(internal.GenerateTestTraceState()), dest)
}

func TestTraceState_CopyTo(t *testing.T) {
	ms := NewTraceState()
	orig := NewTraceState()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = TraceState(internal.GenerateTestTraceState())
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestTraceState_FromRaw_AsRaw(t *testing.T) {
	ms := NewTraceState()
	assert.Equal(t, "", ms.AsRaw())
	ms.FromRaw("congo=t61rcWkgMzE")
	assert.Equal(t, "congo=t61rcWkgMzE", ms.AsRaw())
}
