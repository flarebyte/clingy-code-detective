package aggregator

import (
	"github.com/Masterminds/semver/v3"
)

// MinVersion returns the lower of two semantic versions.
// If both are equal, it returns either.
func MinVersion(v1Str, v2Str string) (*semver.Version, error) {
	v1, err := semver.NewVersion(v1Str)
	if err != nil {
		return nil, err
	}

	v2, err := semver.NewVersion(v2Str)
	if err != nil {
		return nil, err
	}

	if v1.LessThan(v2) {
		return v1, nil
	}
	return v2, nil
}

// MaxVersion returns the higher of two semantic versions.
// If both are equal, it returns either.
func MaxVersion(v1Str, v2Str string) (*semver.Version, error) {
	v1, err := semver.NewVersion(v1Str)
	if err != nil {
		return nil, err
	}

	v2, err := semver.NewVersion(v2Str)
	if err != nil {
		return nil, err
	}

	if v1.GreaterThan(v2) {
		return v1, nil
	}
	return v2, nil
}
