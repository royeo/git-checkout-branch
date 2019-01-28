package main

import "testing"

func Test_isFlag(t *testing.T) {
	tests := []struct {
		args string
		want bool
	}{
		{"help", false},
		{"-h", true},
	}
	for _, tt := range tests {
		if got := isFlag(tt.args); got != tt.want {
			t.Errorf("isFlag() = %v, want %v", got, tt.want)
		}
	}
}
