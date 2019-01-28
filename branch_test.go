package main

import "testing"

func Test_extractBranch(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"feature", "feature"},
		{"hotfix", "hotfix"},
		{"origin/HEAD -> origin/master", "origin/HEAD"},
	}
	for _, tt := range tests {
		if got := extractBranch(tt.name); got != tt.want {
			t.Errorf("extractBranch() = %v, want %v", got, tt.want)
		}
	}
}
