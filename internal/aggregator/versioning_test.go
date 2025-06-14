package aggregator

import (
	"testing"

	"github.com/Masterminds/semver/v3"
)

func TestMinVersion(t *testing.T) {
	tests := []struct {
		v1, v2   string
		expected string
	}{
		{"1.0.0", "2.0.0", "1.0.0"},
		{"2.0.0", "1.0.0", "1.0.0"},
		{"1.2.3", "1.2.3", "1.2.3"},
		{"1.2.3", "1.2.4", "1.2.3"},
		{"2.0.0-beta", "2.0.0", "2.0.0-beta"}, // pre-release is < release
	}

	for _, tt := range tests {
		minV, err := MinVersion(tt.v1, tt.v2)
		if err != nil {
			t.Errorf("MinVersion(%q, %q) returned error: %v", tt.v1, tt.v2, err)
			continue
		}

		expectedV, _ := semver.NewVersion(tt.expected)
		if !minV.Equal(expectedV) {
			t.Errorf("MinVersion(%q, %q) = %q; want %q", tt.v1, tt.v2, minV, expectedV)
		}
	}
}

func TestMaxVersion(t *testing.T) {
	tests := []struct {
		v1, v2   string
		expected string
	}{
		{"1.0.0", "2.0.0", "2.0.0"},
		{"2.0.0", "1.0.0", "2.0.0"},
		{"1.2.3", "1.2.3", "1.2.3"},
		{"1.2.3", "1.2.4", "1.2.4"},
		{"2.0.0-beta", "2.0.0", "2.0.0"}, // release > pre-release
	}

	for _, tt := range tests {
		maxV, err := MaxVersion(tt.v1, tt.v2)
		if err != nil {
			t.Errorf("MaxVersion(%q, %q) returned error: %v", tt.v1, tt.v2, err)
			continue
		}

		expectedV, _ := semver.NewVersion(tt.expected)
		if !maxV.Equal(expectedV) {
			t.Errorf("MaxVersion(%q, %q) = %q; want %q", tt.v1, tt.v2, maxV, expectedV)
		}
	}
}

func TestInvalidVersions(t *testing.T) {
	_, err := MinVersion("invalid", "1.0.0")
	if err == nil {
		t.Error("Expected error for invalid v1 in MinVersion, got nil")
	}

	_, err = MaxVersion("1.0.0", "invalid")
	if err == nil {
		t.Error("Expected error for invalid v2 in MaxVersion, got nil")
	}
}
