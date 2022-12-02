// Copyright © 2022 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
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

package apiclarityexporter

import (
	"time"

	"github.com/dgraph-io/ristretto"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type TraceDatasetMap struct {
	cache  *ristretto.Cache
	keyTTL time.Duration
}

func NewTraceDatasetMap(size int64, ttl time.Duration) (*TraceDatasetMap, error) {
	config := ristretto.Config{
		NumCounters:        (size / 8) * 10,
		MaxCost:            size,
		BufferItems:        64,
		Metrics:            false,
		IgnoreInternalCost: false,
	}
	if cache, err := ristretto.NewCache(&config); err != nil {
		return nil, err
	} else {
		return &TraceDatasetMap{
			cache:  cache,
			keyTTL: ttl, //time.Duration in nanoseconds
		}, nil
	}
}

func (m *TraceDatasetMap) GetDataset(spanID pcommon.SpanID) (string, bool) {
	var key []byte = spanID[:]
	dataset, found := m.cache.Get(key)
	if !found {
		return "", false
	}
	switch d := dataset.(type) {
	default:
		return "", false
	case string:
		return d, true
	}
}

func (m *TraceDatasetMap) PutDataset(spanID pcommon.SpanID, dataset string) bool {
	if dataset == "" {
		return false
	}
	var key []byte = spanID[:]
	success := m.cache.SetWithTTL(key, dataset, int64(len(dataset)), m.keyTTL)
	if success {
		m.cache.Wait() //wait on insert, not lookup
	}
	return success
}
