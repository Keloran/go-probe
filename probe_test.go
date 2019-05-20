package probe

import (
	"github.com/keloran/go-probe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProbe(t *testing.T) {
	tests := []struct {
		expect probe.Healthy
		err    error
	}{
		{
			expect: probe.Healthy{
				Status: "pass",
			},
			err: nil,
		},
	}

	for _, test := range tests {
		resp, err := probe.Probe()
		assert.Equal(t, test.err, err)
		assert.Equal(t, test.expect, resp)
	}
}
