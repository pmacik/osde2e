package installselectors

import (
	"github.com/Masterminds/semver"
	"github.com/openshift/osde2e/pkg/common/config"
	"github.com/openshift/osde2e/pkg/common/spi"
	"github.com/openshift/osde2e/pkg/common/state"
)

func init() {
	registerSelector(oldestClusterImageSet{})
}

// oldestClusterImageSet will use the oldest image from the list of available versions.
type oldestClusterImageSet struct{}

func (o oldestClusterImageSet) ShouldUse() bool {
	return config.Instance.Cluster.UseOldestClusterImageSetForInstall
}

func (o oldestClusterImageSet) Priority() int {
	return 60
}

func (o oldestClusterImageSet) SelectVersion(versionList *spi.VersionList) (*semver.Version, string, error) {
	versionsWithoutDefault := removeDefaultVersion(versionList.AvailableVersions())
	numVersions := len(versionsWithoutDefault)
	versionType := "oldest version"

	// We don't want to fail entirely if there aren't enough versions. It's valid and perhaps even expected
	// that we d on't have enough versions for a middle cluster image set.
	if numVersions < 2 {
		state.Instance.Cluster.EnoughVersionsForOldestOrMiddleTest = false
		return nil, versionType, nil
	}

	return versionsWithoutDefault[0].Version(), versionType, nil
}
