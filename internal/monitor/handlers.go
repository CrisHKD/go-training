package monitor

import (
	"fmt"
)

// SlackNotifier simula el envio de una alerta a Slack
type SlackNotifier struct {
	LastMessage string
}

// OnHealthChange se ejecuta cuando el monitor notifica un evento.
func (s *SlackNotifier) OnHealthChange(e HealthEvent) {
	message := fmt.Sprintf("ALERT: Node %s changed status to %s",
		e.NodeID, e.Status)

	s.LastMessage = message

	// Simulación de envío
	fmt.Println(message)
}

// AutoScaler simula un sistema de auto-escalado
type AutoScaler struct {
	ActiveInstances int
}

// OnHealthChange incrementa instancias si el nodo está sobrecargado
func (a *AutoScaler) OnHealthChange(e HealthEvent) {
	if e.Status == "Overloaded" {
		a.ActiveInstances++
	}
}
