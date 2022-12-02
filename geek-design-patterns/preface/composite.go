package preface

// Flyable .
type Flyable interface {
	Fly()
}

// Tweetable .
type Tweetable interface {
	Tweet()
}

// EggLayable .
type EggLayable interface {
	LayEgg()
}

// Ostrich .
type Ostrich struct{}

// Tweet of Ostrich
func (o *Ostrich) Tweet() {}

// LayEgg of Ostrich
func (o *Ostrich) LayEgg() {}

// Sparrow .
type Sparrow struct{}

// Fly of Sparrow
func (s *Sparrow) Fly() {}

// Tweet of Sparrow
func (s *Sparrow) Tweet() {}

// LayEgg of Sparrow
func (s *Sparrow) LayEgg() {}
