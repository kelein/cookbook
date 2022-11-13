package facade

import "log"

// TV .
type TV struct{}

// ON open TV
func (t *TV) ON() {
	log.Print("OPEN TV")
}

// OFF shutdown TV
func (t *TV) OFF() {
	log.Print("OFF TV")
}

// SoundBox .
type SoundBox struct{}

// ON open SoundBox
func (s *SoundBox) ON() {
	log.Print("OPEN SoundBox")
}

// OFF shutdown SoundBox
func (s *SoundBox) OFF() {
	log.Print("OFF SoundBox")
}

// Light .
type Light struct{}

// ON open Light
func (l *Light) ON() {
	log.Print("OPEN Light")
}

// OFF shutdown Light
func (l *Light) OFF() {
	log.Print("OFF Light")
}

// Xbox .
type Xbox struct{}

// ON open Xbox
func (x *Xbox) ON() {
	log.Print("OPEN Xbox")
}

// OFF shutdown Xbox
func (x *Xbox) OFF() {
	log.Print("OFF Xbox")
}

// MicroPhone .
type MicroPhone struct{}

// ON open MicroPhone
func (m *MicroPhone) ON() {
	log.Print("OPEN MicroPhone")
}

// OFF shutdown	MicroPhone
func (m *MicroPhone) OFF() {
	log.Print("OFF MicroPhone")
}

// Projector .
type Projector struct{}

// ON open Projector
func (p *Projector) ON() {
	log.Print("OPEN Projector")
}

// OFF shutdown Projector
func (p *Projector) OFF() {
	log.Print("OFF Projector")
}

// HomePlayerFacade for family
type HomePlayerFacade struct {
	tv    TV
	light Light
	sound SoundBox
	xbox  Xbox
	mic   MicroPhone
	proj  Projector
}

// KTVMode enter KTV Mode
func (hp *HomePlayerFacade) KTVMode() {
	log.Print("[HomePlayer enter KTV Mode]")
	hp.tv.ON()
	hp.mic.ON()
	hp.proj.ON()
	hp.sound.ON()
	hp.light.OFF()
}

// GameMode enter Game Mode
func (hp *HomePlayerFacade) GameMode() {
	log.Print("[HomePlayer enter Game Mode]")
	hp.tv.ON()
	hp.xbox.ON()
	hp.light.ON()
}
