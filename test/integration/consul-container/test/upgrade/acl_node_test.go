package upgrade

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	libassert "github.com/hashicorp/consul/test/integration/consul-container/libs/assert"
	libcluster "github.com/hashicorp/consul/test/integration/consul-container/libs/cluster"
	libtopology "github.com/hashicorp/consul/test/integration/consul-container/libs/topology"
	"github.com/hashicorp/consul/test/integration/consul-container/libs/utils"
)

// TestACL_Upgrade_Node_Token test verifies the following after upgrade
// 1. Upgraded agent can inherit the persistend token and join the cluster
// 2. Agent token prior to upgrade is still valid after upgrade
func TestACL_Upgrade_Node_Token(t *testing.T) {
	t.Parallel()

	type testcase struct {
		oldversion    string
		targetVersion string
	}
	tcs := []testcase{
		{
			oldversion:    "1.13",
			targetVersion: utils.TargetVersion,
		},
		{
			oldversion:    "1.14",
			targetVersion: utils.TargetVersion,
		},
	}

	run := func(t *testing.T, tc testcase) {
		// NOTE: Disable auto.encrypt due to its conflict with ACL token during bootstrap
		cluster, _, _ := libtopology.NewCluster(t, &libtopology.ClusterConfig{
			NumServers: 1,
			NumClients: 1,
			BuildOpts: &libcluster.BuildOptions{
				Datacenter:           "dc1",
				ConsulVersion:        tc.oldversion,
				InjectAutoEncryption: false,
				ACLEnabled:           true,
			},
			ApplyDefaultProxySettings: true,
		})

		agentToken, err := cluster.CreateAgentToken("dc1",
			cluster.Agents[1].GetAgentName())
		require.NoError(t, err)

		err = cluster.StandardUpgrade(t, context.Background(), tc.targetVersion)
		require.NoError(t, err)

		// Post upgrade validation: agent token can be used to query the node
		// Assert consul catalog nodes  -token e3dc19d9-658d-a430-bcf4-7302efa397fc
		client, err := cluster.Agents[1].NewClient(agentToken, false)
		require.NoError(t, err)
		libassert.CatalogNodeExists(t, client, cluster.Agents[1].GetAgentName())
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("upgrade from %s to %s", tc.oldversion, tc.targetVersion),
			func(t *testing.T) {
				run(t, tc)
			})
	}
}
