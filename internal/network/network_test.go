package network

import "testing"

// ProcessTraffic solo depende del data plane(reenviar paquetes)
func ProcessTraffic(f PacketForwarder) error {
	pkt := []byte("PING")
	return f.Forward(pkt)
}

func TestForwardingOnly(t *testing.T) {
	// LegacySwitch solo implementa PacketForwarder
	var f PacketForwarder = &LegacySwitch{}

	if err := ProcessTraffic(f); err != nil {
		t.Fatalf("se esperaba que el switch reenviara el paquete sin error, pero falló: %v", err)
	}

	// Esto NO compila Porque f es PacketForwarder y esa interfaz NO tiene AddRoute
	// f.AddRoute("10.0.0.0/24", "192.168.1.1")
}

type monitor struct {
	f PacketForwarder
}

func (m monitor) Check(packet []byte) error {
	return m.f.Forward(packet)
}

type controlPanel struct {
	rm RouteManager
}

func (c controlPanel) ProvisionRoute(dest, nextHop string) {
	c.rm.AddRoute(dest, nextHop)
}

func TestCompositeRouter(t *testing.T) {
	r := NewSmartRouter() // SmartRouter implementa ambas interfaces

	// En un sistema que solo necesita PacketForwarder
	mon := monitor{f: r}
	if err := mon.Check([]byte("HELLO")); err != nil {
		t.Fatalf("se esperaba que el router reenviara el paquete sin error, pero falló: %v", err)
	}

	// En un panel que solo necesita RouteManager
	panel := controlPanel{rm: r}
	panel.ProvisionRoute("10.0.0.0/24", "192.168.1.1")

	// Validación simple, la ruta quedo guardada
	if got := r.routes["10.0.0.0/24"]; got != "192.168.1.1" {
		t.Fatalf("ruta mal guardada, se esperaba %q, got %q", "192.168.1.1", got)
	}
}