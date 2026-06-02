package recommend

import (
	"math/rand"
	"reflect"
	"testing"

	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/api/features"
)

func TestSortConditionalUpdatesBySemanticVersions(t *testing.T) {
	expected := []configv1.ConditionalUpdate{
		{Release: configv1.Release{Version: "10.0.0"}},
		{Release: configv1.Release{Version: "2.0.10"}},
		{Release: configv1.Release{Version: "2.0.5"}},
		{Release: configv1.Release{Version: "2.0.1"}},
		{Release: configv1.Release{Version: "2.0.0"}},
		{Release: configv1.Release{Version: "not-sem-ver-2"}},
		{Release: configv1.Release{Version: "not-sem-ver-1"}},
	}

	actual := make([]configv1.ConditionalUpdate, len(expected))
	for i, j := range rand.Perm(len(expected)) {
		actual[i] = expected[j]
	}

	sortConditionalUpdatesBySemanticVersions(actual)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v != %v", actual, expected)
	}
}

func TestIsAcceptRisksEnabled(t *testing.T) {
	featureGate := &configv1.FeatureGate{
		Status: configv1.FeatureGateStatus{
			FeatureGates: []configv1.FeatureGateDetails{
				{
					Version: "4.22.0",
					Enabled: []configv1.FeatureGateAttributes{
						{
							Name: features.FeatureGateClusterUpdateAcceptRisks,
						},
					},
				},
			},
		},
	}

	result := isAcceptRisksEnabled(featureGate, "4.22.0")
	if result != true {
		t.Errorf("ClusterUpdateAcceptRisks feature gate should report as enabled")
	}
}

func TestIsHypershiftEnabled(t *testing.T) {
	infra := &configv1.Infrastructure{
		Status: configv1.InfrastructureStatus{
			ControlPlaneTopology: configv1.ExternalTopologyMode,
		},
	}

	result := isHypershiftEnabled(infra)
	if result != true {
		t.Errorf("Hypershift should report as enabled")
	}
}
