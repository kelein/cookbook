package facade

import "testing"

func TestHomePlayerFacade(t *testing.T) {
	t.Run("HomePlayer", func(t *testing.T) {
		player := new(HomePlayerFacade)
		player.KTVMode()
		player.GameMode()
	})
}

func TestHomePlayerFacade_ON(t *testing.T) {
	t.Run("HomePlayer", func(t *testing.T) {
		player := new(HomePlayerFacade)
		player.tv.ON()
		player.mic.ON()
		player.light.ON()
		player.sound.ON()
		player.xbox.ON()
		player.mic.ON()
		player.proj.ON()
	})
}

func TestHomePlayerFacade_OFF(t *testing.T) {
	t.Run("HomePlayer", func(t *testing.T) {
		player := new(HomePlayerFacade)
		player.tv.OFF()
		player.mic.OFF()
		player.light.OFF()
		player.sound.OFF()
		player.xbox.OFF()
		player.mic.OFF()
		player.proj.OFF()
	})
}
