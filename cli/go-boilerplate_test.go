package main

import (
	"testing"

	"github.com/spf13/cobra"
)

func Test_rootCommand(t *testing.T) {
	cmd := rootCommand()

	tests := []struct {
		name     string
		test     func(cmd *cobra.Command) string
		expected string
	}{
		{
			name:     "Use name",
			test:     func(cmd *cobra.Command) string { return cmd.Use },
			expected: "go-boilerplate",
		},
		{
			name:     "short version",
			test:     func(cmd *cobra.Command) string { return cmd.Short },
			expected: "go-boilerplate is such amaze",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := tt.test(cmd)
			if out != tt.expected {
				t.Errorf("got %q, expected %q", out, tt.expected)
			}
		})
	}
}
