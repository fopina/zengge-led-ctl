package cmd

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanCommand(t *testing.T) {
	// TODO: mock go-ble to be able to test anything at all
	testCases := []struct {
		name     string
		args     []string
		expected string
		err      bool
	}{
		{
			name:     "placeholder",
			args:     []string{"--invalid"},
			expected: "",
			err:      true,
		},
	}

	for _, tc := range testCases {
		cmd := newScanCmd()
		b := bytes.NewBufferString("")

		cmd.SetArgs(tc.args)
		cmd.SetOut(b)

		err := cmd.Execute()
		out, _ := io.ReadAll(b)

		if tc.err {
			assert.Error(t, err, tc.name)
		} else {
			assert.Equal(t, tc.expected, string(out), tc.name)
		}
	}
}
