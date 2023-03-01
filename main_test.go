package main

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"TestMain"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// main()
			assert.Equal(t, "Test Main", "Test Main")
		})
	}
}
