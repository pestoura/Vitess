/*
Copyright 2021 The Vitess Authors.

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

package loadkeyspace

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/stretchr/testify/require"

	"vitess.io/vitess/go/test/endtoend/cluster"
)

var (
	clusterInstance *cluster.LocalProcessCluster
	hostname        = "localhost"
	keyspaceName    = "ks"
	cell            = "zone1"
	sqlSchema       = `
		create table vt_user (
			id bigint,
			name varchar(64),
			primary key (id)
		) Engine=InnoDB;
			
		create table main (
			id bigint,
			val varchar(128),
			primary key(id)
		) Engine=InnoDB;

		create table test_table (
			id bigint,
			val varchar(128),
			primary key(id)
		) Engine=InnoDB;
`
)

func TestBlockedLoadKeyspace(t *testing.T) {
	defer cluster.PanicHandler(t)
	var err error

	clusterInstance = cluster.NewCluster(cell, hostname)
	defer clusterInstance.Teardown()

	// Start topo server
	err = clusterInstance.StartTopo()
	require.NoError(t, err)

	// Start keyspace without the -queryserver-config-schema-change-signal flag
	keyspace := &cluster.Keyspace{
		Name:      keyspaceName,
		SchemaSQL: sqlSchema,
	}
	err = clusterInstance.StartUnshardedKeyspace(*keyspace, 1, false)
	require.NoError(t, err)

	// Start vtgate with the schema_change_signal flag
	clusterInstance.VtGateExtraArgs = []string{"-schema_change_signal"}
	err = clusterInstance.StartVtgate()
	require.Error(t, err)
	logDir := clusterInstance.VtgateProcess.LogDir
	clusterInstance.VtgateProcess = cluster.VtgateProcess{}

	// check fatal logs
	all, err := ioutil.ReadFile(path.Join(logDir, "vtgate.FATAL"))
	require.NoError(t, err)
	require.Contains(t, string(all), "Unable to get initial schema reload")
}
