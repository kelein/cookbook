package adapter

import "testing"

func TestNewAdapter(t *testing.T) {
	t.Run("A", func(t *testing.T) {
		adapter := NewAdapter(new(V220))
		iphone := NewPhone(adapter)
		iphone.Charge()
	})
}
