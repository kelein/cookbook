package finite

import (
	"context"
	"log/slog"
	"testing"
)

func TestDoor_CurrentState(t *testing.T) {
	door := NewDoor("main")

	if door.CurrentState() != StateClosed {
		t.Errorf("Expected state to be %s, got %s", StateClosed, door.CurrentState())
	}

	door.HandleEvent(EventOpen)
	slog.Info("handle event open", "state", door.CurrentState())
	if door.CurrentState() != StateOpen {
		t.Errorf("Expected state to be %s, got %s", StateOpen, door.CurrentState())
	}

	door.HandleEvent(EventClose)
	slog.Info("handle event close", "state", door.CurrentState())
	if door.CurrentState() != StateClosed {
		t.Errorf("Expected state to be %s, got %s", StateClosed, door.CurrentState())
	}
}

func TestNewDoorAdvanced(t *testing.T) {
	door := NewDoorAdvanced("unknown")

	err := door.fsm.Event(context.Background(), EventOpen.Name())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = door.fsm.Event(context.Background(), EventClose.Name())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
