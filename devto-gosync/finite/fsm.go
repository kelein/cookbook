package finite

import (
	"context"
	"log/slog"

	"github.com/fatih/color"
	"github.com/looplab/fsm"
)

// State Enumration for Finite State Machine (FSM)
const (
	StateOpen   State = "open"
	StateClosed State = "closed"
)

// Event Enumration for Finite State Machine (FSM)
const (
	EventOpen  Event = "open"
	EventClose Event = "close"
)

// State for Finite State Machine (FSM)
type State string

// Name returns the name of the state
func (s State) Name() string { return string(s) }

// Event for Finite State Machine (FSM)
type Event string

// Name return the name of the event
func (e Event) Name() string { return string(e) }

// Door represents a door in the FSM
type Door struct {
	to    string
	state State
}

// NewDoor create a initial Door
func NewDoor(to string) *Door {
	return &Door{
		to:    to,
		state: StateClosed,
	}
}

// CurrentState returns the current state of the door
func (d *Door) CurrentState() State { return d.state }

// HandleEvent handles events to change the state of the door
func (d *Door) HandleEvent(event Event) {
	switch event {
	case EventOpen:
		d.state = StateOpen
	case EventClose:
		d.state = StateClosed
	}
}

// DoorAdvanced represents an advanced door with a finite state machine
type DoorAdvanced struct {
	to  string
	fsm *fsm.FSM
}

// NewDoorAdvanced creates a new advanced door with a finite state machine
func NewDoorAdvanced(to string) *DoorAdvanced {
	door := &DoorAdvanced{to: to}
	door.fsm = fsm.NewFSM(
		StateClosed.Name(),
		fsm.Events{
			{
				Name: EventOpen.Name(),
				Src:  []string{StateClosed.Name()},
				Dst:  StateOpen.Name(),
			},
			{
				Name: EventClose.Name(),
				Src:  []string{StateOpen.Name()},
				Dst:  StateClosed.Name(),
			},
		},
		fsm.Callbacks{
			"onEvent": func(ctx context.Context, e *fsm.Event) {
				door.onEvent(e)
			},

			"before_open": func(_ context.Context, e *fsm.Event) {
				color.Magenta("| before open\t | %s | %s |", e.Src, e.Dst)
			},

			// 任一事件发生之前触发
			"before_event": func(_ context.Context, e *fsm.Event) {
				color.HiMagenta("| before event\t | %s | %s |", e.Src, e.Dst)
			},

			// 在离开 closed 状态时触发
			"leave_closed": func(_ context.Context, e *fsm.Event) {
				color.Cyan("| leave closed\t | %s | %s |", e.Src, e.Dst)
			},

			// 离开任一状态时触发
			"leave_state": func(_ context.Context, e *fsm.Event) {
				color.HiCyan("| leave state\t | %s | %s |", e.Src, e.Dst)
			},

			// 在进入 open 状态时触发（这里的 open 是指代 open state）
			"enter_open": func(_ context.Context, e *fsm.Event) {
				color.Green("| enter open\t | %s | %s |", e.Src, e.Dst)
			},

			// 进入任一状态时触发
			"enter_state": func(_ context.Context, e *fsm.Event) {
				color.HiGreen("| enter state\t | %s | %s |", e.Src, e.Dst)
			},

			// 在 open 事件发生之后触发（这里的 open 是指代 open event）
			"after_open": func(_ context.Context, e *fsm.Event) {
				color.Yellow("| after open\t | %s | %s |", e.Src, e.Dst)
			},

			// 任一事件结束后触发
			"after_event": func(_ context.Context, e *fsm.Event) {
				color.HiYellow("| after event\t | %s | %s |", e.Src, e.Dst)
			},
		})

	return door
}

func (d *DoorAdvanced) onEvent(e *fsm.Event) {
	slog.Info("on event", "name", e.Event, "from", e.Src, "to", e.Dst)
}
