package observer

import (
	"fmt"
	"log"
)

const maxNotifyCount = 20

// Listener of abstract Event
type Listener interface {
	Title() string
	GetName() string
	GetParty() string
	OnBeFighted(event *Event)
}

// Notifier of abstract Event
type Notifier interface {
	Attach(listener Listener)
	Detach(listener Listener)
	Notify(event *Event)
}

// Event of concrete
type Event struct {
	Owner    Listener
	Receiver Listener
	Notifier Notifier
	Message  string
}

// Hero stands for concrete Listener
type Hero struct {
	Name  string
	Party string
}

// Title of Hero
func (h *Hero) Title() string {
	return fmt.Sprintf("[%s]%s", h.Party, h.Name)
}

// GetName return the name of the hero
func (h *Hero) GetName() string { return h.Name }

// GetParty return the party of the hero
func (h *Hero) GetParty() string { return h.Party }

// Fight metheod of the hero
func (h *Hero) Fight(receiver Listener, n Notifier) {
	msg := fmt.Sprintf("%s Fight %s ...", h.Title(), receiver.Title())

	event := &Event{
		Owner:    h,
		Notifier: n,
		Message:  msg,
		Receiver: receiver,
	}
	n.Notify(event)
}

func (h *Hero) related(e *Event) bool {
	return h.Name == e.Owner.GetName() ||
		h.Name == e.Receiver.GetName()
}

// OnBeFighted metheod of the hero
func (h *Hero) OnBeFighted(event *Event) {
	if h.related(event) {
		return
	}

	if h.Party == event.Owner.GetParty() {
		log.Printf("%s Clap Hands and Say Yeah!", h.Title())
		return
	}

	if h.Party == event.Receiver.GetParty() {
		log.Printf("%s Keep Fighting!", h.Title())
		h.Fight(event.Owner, event.Notifier)
		return
	}
}

// Informer stands for concrete Listener
type Informer struct {
	heros []Listener
	count int
}

// Attach add listener into Informer list
func (in *Informer) Attach(listener Listener) {
	in.heros = append(in.heros, listener)
}

// Detach remove listener from Informer list
func (in *Informer) Detach(listener Listener) {
	for idx, l := range in.heros {
		if listener == l {
			in.heros = append(in.heros[:idx], in.heros[idx+1:]...)
			break
		}
	}
}

// Notify send notification to each Listener
func (in *Informer) Notify(event *Event) {
	log.Printf("[Informer] current Notify count: %v", in.count)
	if in.count >= maxNotifyCount {
		log.Print("[Informer] Notify Limited")
		return
	}

	log.Printf("[Informer] Notification: %s", event.Message)
	in.count++
	for _, hero := range in.heros {
		hero.OnBeFighted(event)
	}
}
