package probe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Healthy structure
type Healthy struct {
	Status string `json:"status"`
}

// HTTP the response itself
func HTTP(w http.ResponseWriter, r *http.Request) {
	// request is by a human
	buf, _ := ioutil.ReadAll(r.Body)
	if len(buf) >= 1 {
		fmt.Printf("probe request: %s\n", string(buf))
	}

	// get response
	resp, _ := Probe()

	// send status
	j, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/health+json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(j)
	if err != nil {
		fmt.Printf("write failed: %v\n", err)
	}
}

// Probe do the response
func Probe() (Healthy, error) {
	return Healthy{
		Status: "pass",
	}, nil
}
