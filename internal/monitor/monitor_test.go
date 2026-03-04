package monitor

import (
	"testing"
	"time"
)

// SpyObserver es un observer espía para tests solo cuenta cuantas veces fue llamado y que evento recibio
type SpyObserver struct {
	Calls int
	Last  HealthEvent
}

func (s *SpyObserver) OnHealthChange(e HealthEvent) {
	s.Calls++
	s.Last = e
}

func TestObserverRegistration(t *testing.T) {
	monitor := NewClusterMonitor()

	// Usamos observers
	slack := &SlackNotifier{}
	auto := &AutoScaler{}

	monitor.Register(slack)
	monitor.Register(auto)

	// Como el test esta en el mismo package, puede ver el campo privado observers
	if len(monitor.observers) != 2 {
		t.Fatalf("se esperaban %d observers registrados, got %d", 2, len(monitor.observers))
	}
}

func TestNotificationFlow(t *testing.T) {
	monitor := NewClusterMonitor()

	// Creamos 2 spies para verificar que Notify llama a todos
	spy1 := &SpyObserver{}
	spy2 := &SpyObserver{}

	monitor.Register(spy1)
	monitor.Register(spy2)

	event := HealthEvent{
		NodeID:    "node-1",
		Status:    "Down",
		Timestamp: time.Date(2026, 3, 3, 12, 0, 0, 0, time.UTC),
	}

	monitor.Notify(event)

	// Verificamos que ambos recibieron 1 llamada.
	if spy1.Calls != 1 {
		t.Fatalf("spy1: se esperaba %d llamada, got %d", 1, spy1.Calls)
	}
	if spy2.Calls != 1 {
		t.Fatalf("spy2: se esperaba %d llamada, got %d", 1, spy2.Calls)
	}

	// verificamos que el evento que recibieron es el mismo.
	if spy1.Last != event {
		t.Fatalf("spy1: evento recibido distinto al esperado")
	}
	if spy2.Last != event {
		t.Fatalf("spy2: evento recibido distinto al esperado")
	}
}

func TestFullAlertPipeline(t *testing.T) {
	monitor := NewClusterMonitor()

	// Observers del ejercicio
	slack := &SlackNotifier{}
	auto := &AutoScaler{}

	monitor.Register(slack)
	monitor.Register(auto)

	event := HealthEvent{
		NodeID:    "node-99",
		Status:    "Overloaded", // este estado debe disparar el autoscaling
		Timestamp: time.Date(2026, 3, 3, 12, 30, 0, 0, time.UTC),
	}

	monitor.Notify(event)

	// Valida que AutoScaler reacciono
	if auto.ActiveInstances != 1 {
		t.Fatalf("se esperaba ActiveInstances=%d, got %d", 1, auto.ActiveInstances)
	}

	// Valida que SlackNotifier genero el mensaje correcto
	expectedMsg := "ALERT: Node node-99 changed status to Overloaded"
	if slack.LastMessage != expectedMsg {
		t.Fatalf("mensaje Slack incorrecto.\nEsperado: %q\nGot:      %q", expectedMsg, slack.LastMessage)
	}
}
