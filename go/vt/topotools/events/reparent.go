/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package events defines the structures used for events dispatched from the
// wrangler package.
package events

import (
	base "vitess.io/vitess/go/vt/events"
	"vitess.io/vitess/go/vt/topo"

	topodatapb "vitess.io/vitess/go/vt/proto/topodata"
)

// Reparent is an event that describes a single step in the reparent process.
type Reparent struct {
	base.StatusUpdater

	ShardInfo              topo.ShardInfo
	OldPrimary, NewPrimary *topodatapb.Tablet
	ExternalID             string
}
