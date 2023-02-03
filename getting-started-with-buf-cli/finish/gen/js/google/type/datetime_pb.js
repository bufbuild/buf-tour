// Copyright 2021 Google LLC
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

// @generated by protoc-gen-es v1.0.0 with parameter "target=js"
// @generated from file google/type/datetime.proto (package google.type, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Duration, proto3 } from "@bufbuild/protobuf";

/**
 * Represents civil time (or occasionally physical time).
 *
 * This type can represent a civil time in one of a few possible ways:
 *
 *  * When utc_offset is set and time_zone is unset: a civil time on a calendar
 *    day with a particular offset from UTC.
 *  * When time_zone is set and utc_offset is unset: a civil time on a calendar
 *    day in a particular time zone.
 *  * When neither time_zone nor utc_offset is set: a civil time on a calendar
 *    day in local time.
 *
 * The date is relative to the Proleptic Gregorian Calendar.
 *
 * If year is 0, the DateTime is considered not to have a specific year. month
 * and day must have valid, non-zero values.
 *
 * This type may also be used to represent a physical time if all the date and
 * time fields are set and either case of the `time_offset` oneof is set.
 * Consider using `Timestamp` message for physical time instead. If your use
 * case also would like to store the user's timezone, that can be done in
 * another field.
 *
 * This type is more flexible than some applications may want. Make sure to
 * document and validate your application's limitations.
 *
 * @generated from message google.type.DateTime
 */
export const DateTime = proto3.makeMessageType(
  "google.type.DateTime",
  () => [
    { no: 1, name: "year", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "month", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "day", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "hours", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "minutes", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "seconds", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "nanos", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 8, name: "utc_offset", kind: "message", T: Duration, oneof: "time_offset" },
    { no: 9, name: "time_zone", kind: "message", T: TimeZone, oneof: "time_offset" },
  ],
);

/**
 * Represents a time zone from the
 * [IANA Time Zone Database](https://www.iana.org/time-zones).
 *
 * @generated from message google.type.TimeZone
 */
export const TimeZone = proto3.makeMessageType(
  "google.type.TimeZone",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

