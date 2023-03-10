// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestScopeMetrics_MoveTo(t *testing.T) {
	ms := generateTestScopeMetrics()
	dest := NewScopeMetrics()
	ms.MoveTo(dest)
	assert.Equal(t, NewScopeMetrics(), ms)
	assert.Equal(t, generateTestScopeMetrics(), dest)
}

func TestScopeMetrics_CopyTo(t *testing.T) {
	ms := NewScopeMetrics()
	orig := NewScopeMetrics()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestScopeMetrics()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestScopeMetrics_Scope(t *testing.T) {
	ms := NewScopeMetrics()
	internal.FillTestInstrumentationScope(internal.InstrumentationScope(ms.Scope()))
	assert.Equal(t, pcommon.InstrumentationScope(internal.GenerateTestInstrumentationScope()), ms.Scope())
}

func TestScopeMetrics_SchemaUrl(t *testing.T) {
	ms := NewScopeMetrics()
	assert.Equal(t, "", ms.SchemaUrl())
	ms.SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	assert.Equal(t, "https://opentelemetry.io/schemas/1.5.0", ms.SchemaUrl())
}

func TestScopeMetrics_Metrics(t *testing.T) {
	ms := NewScopeMetrics()
	assert.Equal(t, NewMetricSlice(), ms.Metrics())
	fillTestMetricSlice(ms.Metrics())
	assert.Equal(t, generateTestMetricSlice(), ms.Metrics())
}

func generateTestScopeMetrics() ScopeMetrics {
	tv := NewScopeMetrics()
	fillTestScopeMetrics(tv)
	return tv
}

func fillTestScopeMetrics(tv ScopeMetrics) {
	internal.FillTestInstrumentationScope(internal.NewInstrumentationScope(&tv.orig.Scope))
	tv.orig.SchemaUrl = "https://opentelemetry.io/schemas/1.5.0"
	fillTestMetricSlice(newMetricSlice(&tv.orig.Metrics))
}
