package probe_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/keloran/go-probe"
	"github.com/stretchr/testify/assert"
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

func TestHTTP(t *testing.T) {
	tests := []struct {
		expect probe.Healthy
	}{
		{
			expect: probe.Healthy{
				Status: "pass",
			},
		},
	}

	for _, test := range tests {
		request, _ := http.NewRequest("GET", "/", bytes.NewBuffer([]byte{}))
		response := httptest.NewRecorder()
		probe.HTTP(response, request)
		assert.Equal(t, 200, response.Code)
		body, _ := ioutil.ReadAll(response.Body)
		healthy := probe.Healthy{}
		_ = json.Unmarshal(body, &healthy)
		assert.Equal(t, test.expect, healthy)
	}
}
