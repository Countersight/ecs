// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by scripts/gocodegen.go - DO NOT EDIT.

package ecs

import (
	"time"
)

// The Countersight fields describe events that have occurred during the ingest
// into the Countersight system.
type Countersight struct {
	// The name of the pipeline that received the event.
	Pipeline string `ecs:"pipeline"`

	// The name and instance of collector that received the event.
	CollectorName string `ecs:"collector.name"`

	// The type of collector that received the event.
	CollectorType string `ecs:"collector.type"`

	// Trace contains a list of processors that have seen this event during
	// ingestion.
	Trace string `ecs:"trace"`

	// The date and time that this event was processed by the Countersight
	// system
	IngestTime time.Time `ecs:"ingest_time"`

	// The countersight.tags field contains processing messages, such as when
	// an event  failed to be properly parsed.
	Tags string `ecs:"tags"`
}
