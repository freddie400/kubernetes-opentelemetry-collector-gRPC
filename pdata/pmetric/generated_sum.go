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
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
)

// Sum represents the type of a numeric metric that is calculated as a sum of all reported measurements over a time interval.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSum function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Sum struct {
	orig *otlpmetrics.Sum
}

func newSum(orig *otlpmetrics.Sum) Sum {
	return Sum{orig}
}

// NewSum creates a new empty Sum.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewSum() Sum {
	return newSum(&otlpmetrics.Sum{})
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms Sum) MoveTo(dest Sum) {
	*dest.orig = *ms.orig
	*ms.orig = otlpmetrics.Sum{}
}

// AggregationTemporality returns the aggregationtemporality associated with this Sum.
func (ms Sum) AggregationTemporality() AggregationTemporality {
	return AggregationTemporality(ms.orig.AggregationTemporality)
}

// SetAggregationTemporality replaces the aggregationtemporality associated with this Sum.
func (ms Sum) SetAggregationTemporality(v AggregationTemporality) {
	ms.orig.AggregationTemporality = otlpmetrics.AggregationTemporality(v)
}

// IsMonotonic returns the ismonotonic associated with this Sum.
func (ms Sum) IsMonotonic() bool {
	return ms.orig.IsMonotonic
}

// SetIsMonotonic replaces the ismonotonic associated with this Sum.
func (ms Sum) SetIsMonotonic(v bool) {
	ms.orig.IsMonotonic = v
}

// DataPoints returns the DataPoints associated with this Sum.
func (ms Sum) DataPoints() NumberDataPointSlice {
	return newNumberDataPointSlice(&ms.orig.DataPoints)
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms Sum) CopyTo(dest Sum) {
	dest.SetAggregationTemporality(ms.AggregationTemporality())
	dest.SetIsMonotonic(ms.IsMonotonic())
	ms.DataPoints().CopyTo(dest.DataPoints())
}
