package monitor

import "time"

// HealthEvent representa un cambio de salud
// Es lo que el monitor envia a todos los observers
type HealthEvent struct {
	NodeID    string
	Status    string
	Timestamp time.Time
}

// HealthObserver es el contrato que deben cumplir los reaccionadores
type HealthObserver interface {
	OnHealthChange(event HealthEvent)
}
