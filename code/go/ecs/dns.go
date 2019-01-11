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

// DNS-specific event fields. This is NOT YET PART OF THE OFFICIAL ECS This
// schema has been taken from an open PR for DNS fields against the ECS  See
// https://github.com/elastic/ecs/pull/180  Fingers crossed it makes it in
// unmodified  The ecs does not support keyword values, so in the interim, I'll
// change the keywords (flags)  To be keyword, where we can add 'true' or
// 'false'  I had to make a few modifications from the PR that guy created.
//   - removed the phase entries
//   - added a type to a bunch of entries that were missing types
//   - removed all of the 'type: object' entries
type Dns struct {
	// The DNS packet identifier assigned by the program that generated the
	// query. The identifier is copied to the response.
	ID string `ecs:"id"`

	// The DNS operation code that specifies the kind of query in the message.
	// This value is set by the originator of a query and copied into the
	// response.
	OpCode string `ecs:"op_code"`

	// A DNS flag specifying that the responding server is an authority for the
	// domain name used in the question.
	FlagsAuthoritative string `ecs:"flags.authoritative"`

	// A DNS flag specifying whether recursive query support is available in
	// the name server.
	FlagsRecursionAvailable string `ecs:"flags.recursion_available"`

	// A DNS flag specifying that the client directs the server to pursue a
	// query recursively. Recursive query support is optional.
	FlagsRecursionDesired string `ecs:"flags.recursion_desired"`

	// A DNS flag specifying that the recursive server considers the response
	// authentic.
	FlagsAuthenticData string `ecs:"flags.authentic_data"`

	// A DNS flag specifying that the client disables the server signature
	// validation of the query.
	FlagsCheckingDisabled string `ecs:"flags.checking_disabled"`

	// A DNS flag specifying that only the first 512 bytes of the reply were
	// returned.
	FlagsTruncatedResponse string `ecs:"flags.truncated_response"`

	// The DNS status code.
	ResponseCode string `ecs:"response_code"`

	// The domain name being queried. If the name field contains non-printable
	// characters (below 32 or above 126), then those characters are
	// represented as escaped base 10 integers (\DDD). Back slashes and quotes
	// are escaped. Tabs, carriage returns, and line feeds are converted to \t,
	// \r, and \n respectively.
	QuestionName string `ecs:"question.name"`

	// The type of records being queried.
	QuestionType string `ecs:"question.type"`

	// The class of of records being queried.
	QuestionClass string `ecs:"question.class"`

	// The effective top-level domain (eTLD) plus one more label. For example,
	// the eTLD+1 for "foo.bar.golang.org." is "golang.org.". The data for
	// determining the eTLD comes from an embedded copy of the data from
	// http://publicsuffix.org.
	QuestionEtldPlusOne string `ecs:"question.etld_plus_one"`

	// The size of name being queried (in bytes).
	QuestionSize string `ecs:"question.size"`

	// The number of resource records contained in the `dns.answers` field.
	AnswersCount int64 `ecs:"answers_count"`

	// The domain name to which this resource record pertains.
	AnswersName string `ecs:"answers.name"`

	// The type of data contained in this resource record.
	AnswersType string `ecs:"answers.type"`

	// The class of DNS data contained in this resource record.
	AnswersClass string `ecs:"answers.class"`

	// The time interval in seconds that this resource record may be cached
	// before it should be discarded. Zero values mean that the data should not
	// be cached.
	AnswersTtl int64 `ecs:"answers.ttl"`

	// The data describing the resource. The meaning of this data depends on
	// the type and class of the resource record.
	AnswersData string `ecs:"answers.data"`

	// The number of resource records contained in the `dns.authorities` field.
	AuthoritiesCount int64 `ecs:"authorities_count"`

	// The domain name to which this resource record pertains.
	AuthoritiesName string `ecs:"authorities.name"`

	// The type of data contained in this resource record.
	AuthoritiesType string `ecs:"authorities.type"`

	// The class of DNS data contained in this resource record.
	AuthoritiesClass string `ecs:"authorities.class"`

	// The number of resource records contained in the `dns.additionals` field.
	AdditionalsCount int64 `ecs:"additionals_count"`

	// The domain name to which this resource record pertains.
	AdditionalsName string `ecs:"additionals.name"`

	// The type of data contained in this resource record.
	AdditionalsType string `ecs:"additionals.type"`

	// The class of DNS data contained in this resource record.
	AdditionalsClass string `ecs:"additionals.class"`

	// The time interval in seconds that this resource record may be cached
	// before it should be discarded. Zero values mean that the data should not
	// be cached.
	AdditionalsTtl int64 `ecs:"additionals.ttl"`

	// The data describing the resource. The meaning of this data depends on
	// the type and class of the resource record.
	AdditionalsData string `ecs:"additionals.data"`

	// The EDNS version.
	OptVersion string `ecs:"opt.version"`

	// If set, the transaction uses DNSSEC.
	OptDo string `ecs:"opt.do"`

	// Extended response code field.
	OptExtRcode string `ecs:"opt.ext_rcode"`

	// Requestor's UDP payload size (in bytes).
	OptUdpSize int64 `ecs:"opt.udp_size"`
}
