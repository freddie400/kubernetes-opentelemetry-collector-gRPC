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

	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
)

func TestSum_MoveTo(t *testing.T) {
	ms := generateTestSum()
	dest := NewSum()
	ms.MoveTo(dest)
	assert.Equal(t, NewSum(), ms)
	assert.Equal(t, generateTestSum(), dest)
}

func TestSum_CopyTo(t *testing.T) {
	ms := NewSum()
	orig := NewSum()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestSum()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestSum_AggregationTemporality(t *testing.T) {
	ms := NewSum()
	assert.Equal(t, AggregationTemporality(otlpmetrics.AggregationTemporality(0)), ms.AggregationTemporality())
	testValAggregationTemporality := AggregationTemporality(otlpmetrics.AggregationTemporality(1))
	ms.SetAggregationTemporality(testValAggregationTemporality)
	assert.Equal(t, testValAggregationTemporality, ms.AggregationTemporality())
}

func TestSum_IsMonotonic(t *testing.T) {
	ms := NewSum()
	assert.Equal(t, false, ms.IsMonotonic())
	ms.SetIsMonotonic(true)
	assert.Equal(t, true, ms.IsMonotonic())
}

func TestSum_DataPoints(t *testing.T) {
	ms := NewSum()
	assert.Equal(t, NewNumberDataPointSlice(), ms.DataPoints())
	fillTestNumberDataPointSlice(ms.DataPoints())
	assert.Equal(t, generateTestNumberDataPointSlice(), ms.DataPoints())
}

func generateTestSum() Sum {
	tv := NewSum()
	fillTestSum(tv)
	return tv
}

func fillTestSum(tv Sum) {
	tv.orig.AggregationTemporality = otlpmetrics.AggregationTemporality(1)
	tv.orig.IsMonotonic = true
	fillTestNumberDataPointSlice(newNumberDataPointSlice(&tv.orig.DataPoints))
}
