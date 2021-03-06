package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testNodeAwsCluster = Cluster{
	Name:               "Test AWS Cluster Go SDK " + GetTicks(),
	Provider:           "aws",
	MasterCount:        1,
	MasterSize:         "t2.medium",
	WorkerCount:        2,
	WorkerSize:         "t2.medium",
	Region:             "us-east-1",
	Zone:               "us-east-1a",
	ProviderNetworkID:  "__new__",
	ProviderNetworkCdr: "172.23.0.0/16",
	ProviderSubnetID:   "__new__",
	ProviderSubnetCidr: "172.23.1.0/24",
	KubernetesVersion:  "v1.13.1",
	RbacEnabled:        true,
	DashboardEnabled:   true,
	EtcdType:           "classic",
	Platform:           "coreos",
	Channel:            "stable",
	NetworkComponents:  []NetworkComponent{},
	Solutions:          []Solution{Solution{Solution: "helm_tiller"}},
}

func TestLiveBasicNode(t *testing.T) {
	clusterID := testNodeClusterCreate(t)
	nodeID := testNodeCreate(t, clusterID)
	testNodeList(t, clusterID)
	testNodeGet(t, clusterID, nodeID)
	testNodeDelete(t, clusterID, nodeID)
	testNodeClusterDelete(t, clusterID)
}

func testNodeClusterCreate(t *testing.T) int {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	if err != nil {
		t.Error(err)
	}

	awsKeysetID, err := GetIDFromEnv("NKS_AWS_KEYSET")
	if err != nil {
		t.Error(err)
	}

	testNodeAwsCluster.ProviderKey = awsKeysetID
	testNodeAwsCluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testNodeAwsCluster)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitClusterRunning(orgID, cluster.ID, true, timeout)

	return cluster.ID
}

func testNodeCreate(t *testing.T, clusterID int) int {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	nodeAdd := NodeAdd{
		Count:              1,
		Size:               "t2.medium",
		Role:               "master",
		Zone:               "us-east-1a",
		ProviderSubnetID:   "__new__",
		ProviderSubnetCidr: "172.23.1.0/24",
		RootDiskSize:       50,
	}

	nodes, err := c.AddNode(orgID, clusterID, nodeAdd)
	if err != nil {
		t.Error(err)
	}

	node := nodes[0]

	err = c.WaitNodeProvisioned(orgID, clusterID, node.ID, timeout)
	if err != nil {
		t.Error(err)
	}

	return node.ID
}

func testNodeList(t *testing.T, clusterID int) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	list, err := c.GetNodes(orgID, clusterID)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, len(list), 4, "There should be 4 nodes")
}

func testNodeGet(t *testing.T, clusterID, nodeID int) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	node, err := c.GetNode(orgID, clusterID, nodeID)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, node)
	assert.Equal(t, node.ID, nodeID, "Master node must exist")
}

func testNodeDelete(t *testing.T, clusterID, nodeID int) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = c.DeleteNode(orgID, clusterID, nodeID)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitNodeDeleted(orgID, clusterID, nodeID, timeout)
	if err != nil {
		t.Error(err)
	}
}

func testNodeClusterDelete(t *testing.T, clusterID int) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = c.DeleteCluster(orgID, clusterID)
	if err != nil {
		t.Error(err)
	}

	err = c.WaitClusterDeleted(orgID, clusterID, timeout)
	if err != nil {
		t.Error(err)
	}
}
