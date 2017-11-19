package suzuri

import (
	"regexp"
	"testing"
)

func TestVersion(t *testing.T) {
	versionRegexp := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	if !versionRegexp.MatchString(version) {
		t.Errorf("expected %v, got %v", versionRegexp, version)
	}
}
