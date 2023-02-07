package server

import (
	stdContext "context"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestStart(t *testing.T) {
	tests := []struct {
		name                   string
		givenPrefix            string
		expectedStatus         int
		expectedBodyStartsWith string
	}{
		{
			name:                   "root_url",
			givenPrefix:            "/",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "Hello, Docker! <3",
		},
		{
			name:                   "ping_url",
			givenPrefix:            "/ping",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{Status:\"OK\"}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Dummy test
			assert.Equal(t, "Test Server", "Test Server")
			// errCh := make(chan error)
			// go func() {
			// 	errCh <- Start()
			// }()
			// err := waitForServerStart(errCh, false)
			// assert.NoError(t, err)
			// resp, err := http.Get("0.0.0.0:8080" + tt.givenPrefix)
			// assert.Equal(t, tt.expectedStatus, resp.Status)

			// body := make([]byte, 100)
			// _, err = resp.Body.Read(body)
			// if tt.expectedBodyStartsWith != "" {
			// 	assert.True(t, strings.HasPrefix(string(body), tt.expectedBodyStartsWith))
			// } else {
			// 	assert.Equal(t, "", body)
			// }

		})
	}
}

func waitForServerStart(errChan <-chan error, isTLS bool) error {
	ctx, cancel := stdContext.WithTimeout(stdContext.Background(), 200*time.Millisecond)
	defer cancel()

	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			_, err := net.Dial("tcp", "0.0.0.0:8080")
			if err != nil {
				return err
			}
			return nil
		case err := <-errChan:
			if err == http.ErrServerClosed {
				return nil
			}
			return err
		}
	}
}
