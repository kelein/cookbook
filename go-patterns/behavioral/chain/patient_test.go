package chain

import (
	"testing"
)

func TestPatientHandler_Do(t *testing.T) {
	t.Run("PatientHandler", func(t *testing.T) {
		healthHandler := &StartHandler{}
		healthHandler.SetNext(&Reception{}).
			SetNext(&Clinic{}).
			SetNext(&Cashier{}).
			SetNext(&Pharmacy{})

		patient := &Patient{Name: "Cathy"}
		if err := healthHandler.Execute(patient); err != nil {
			t.Errorf("execute handler error: %v", err)
		}
	})
}
