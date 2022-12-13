package optional

import (
	"testing"
	"time"
)

func TestServer_Start(t *testing.T) {
	t.Run("option", func(t *testing.T) {
		server := NewServer(
			WithHost("127.0.0.1"),
			WithPort(8080),
			WithMaxConn(65535),
			WithTimeout(time.Second*60),
		)

		if err := server.Start(); err != nil {
			t.Errorf("Server.Start() error = %v", err)
		}
	})
}
