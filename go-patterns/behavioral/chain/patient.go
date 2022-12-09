// Package chain for <Chain of Responsibility>
package chain

import "log"

// Patient object
type Patient struct {
	Name         string
	RegisterDone bool
	DiagnoseDone bool
	MedicineDone bool
	PaymentDone  bool
}

// PatientHandler of abstract
type PatientHandler interface {
	SetNext(PatientHandler) PatientHandler
	Execute(*Patient) error
	Do(*Patient) error
}

// Next base handler
type Next struct {
	nextHandler PatientHandler
}

// SetNext sets the next handler and return it
func (n *Next) SetNext(h PatientHandler) PatientHandler {
	n.nextHandler = h
	return h
}

// Execute exec next handlers
func (n *Next) Execute(p *Patient) error {
	if n.nextHandler == nil {
		return nil
	}
	if err := n.nextHandler.Do(p); err != nil {
		return err
	}
	return n.nextHandler.Execute(p)
}

// Reception request
type Reception struct{ Next }

// Do of Reception
func (r *Reception) Do(p *Patient) error {
	if p.RegisterDone {
		log.Printf("[Reception] patient %q already registered", p.Name)
		return nil
	}
	log.Printf("[Reception] patient %q on registering ...", p.Name)
	p.RegisterDone = true
	return nil
}

// Clinic request
type Clinic struct{ Next }

// Do of Clinic
func (c *Clinic) Do(p *Patient) error {
	if p.DiagnoseDone {
		log.Printf("[Clinic] patient %q alreay diagnosed", p.Name)
		return nil
	}
	log.Printf("[Clinic] patient %q on checking ...", p.Name)
	p.DiagnoseDone = true
	return nil
}

// Pharmacy request
type Pharmacy struct{ Next }

// Do of Pharmacy
func (y *Pharmacy) Do(p *Patient) error {
	if p.MedicineDone {
		log.Printf("[Pharmacy] patient %q already got Medicines", p.Name)
		return nil
	}
	log.Printf("[Pharmacy] patient %q on giving Medicines", p.Name)
	p.MedicineDone = true
	return nil
}

// Cashier request
type Cashier struct{ Next }

// Do of Cashier
func (c *Cashier) Do(p *Patient) error {
	if p.PaymentDone {
		log.Printf("[Cashier] patient %q already payed", p.Name)
		return nil
	}
	log.Printf("[Cashier] patient %q on paying ...", p.Name)
	p.PaymentDone = true
	return nil
}

// StartHandler send request to next handler
type StartHandler struct{ Next }

// Do of StartHandler
func (s *StartHandler) Do(p *Patient) error {
	return nil
}
