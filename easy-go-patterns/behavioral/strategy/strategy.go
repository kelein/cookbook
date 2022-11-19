package strategy

import "log"

// WeaponStrategy of abstract
type WeaponStrategy interface {
	Attack()
}

// Rifle stands for concrete weapon
type Rifle struct{}

// Attack of Rifle
func (r *Rifle) Attack() {
	log.Print("Attack by Rifle")
}

// Knife stands for concrete weapon
type Knife struct{}

// Attack of Knife
func (k *Knife) Attack() {
	log.Print("Attack by Knife")
}

// Hero with WeaponStrategy
type Hero struct {
	strategy WeaponStrategy
}

// SetWeaponStrategy of Hero
func (h *Hero) SetWeaponStrategy(w WeaponStrategy) {
	h.strategy = w
}

// Attack of Hero
func (h *Hero) Attack() {
	h.strategy.Attack()
}
