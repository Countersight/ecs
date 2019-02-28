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

// The `base` field set contains all fields which are on the top level. These
// fields are common across all types of events.
type Base struct {
	// Date/time when the event originated.
	// This is the date/time extracted from the event, typically representing
	// when the event was generated by the source.
	// If the event source has no original timestamp, this value is typically
	// populated by the first time the event was received by the pipeline.
	// Required field for all events.
	Timestamp time.Time `ecs:"@timestamp"`

	// List of keywords used to tag each event.
	Tags string `ecs:"tags"`

	// List of filters that processed this event
	ProcessingMessages string `ecs:"processing_messages"`

	// Custom key/value pairs.
	// Can be used to add meta information to events. Should not contain nested
	// objects. All values are stored as keyword.
	// Example: `docker` and `k8s` labels.
	Labels map[string]interface{} `ecs:"labels"`

	// For log events the message field contains the log message, optimized for
	// viewing in a log viewer.
	// For structured logs without an original message field, other fields can
	// be concatenated to form a human-readable summary of the event.
	// If multiple messages exist, they can be combined into one message.
	Message string `ecs:"message"`
}
