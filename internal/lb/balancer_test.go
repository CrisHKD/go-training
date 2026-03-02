package lb

import "testing"

func TestRoundRobin(t *testing.T) {
	lb := NewLoadBalancer(&RoundRobinStrategy{})

	lb.AddNode(Node{ID: "A", Address: "10.0.0.1:8080"})
	lb.AddNode(Node{ID: "B", Address: "10.0.0.2:8080"})
	lb.AddNode(Node{ID: "C", Address: "10.0.0.3:8080"})

	want := []string{"A", "B", "C", "A"}
	got := make([]string, 0, len(want))

	for i := 0; i < len(want); i++ {
		n := lb.GetNext()
		if n == nil {
			t.Fatalf("se esperaba un nodo, pero GetNext() devolvió nil en la llamada %d", i)
		}
		got = append(got, n.ID)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("round robin incorrecto en la posición %d: esperado %q, obtenido %q", i, want[i], got[i])
		}
	}
}

func TestLeastConn(t *testing.T) {
	lb := NewLoadBalancer(&LeastConnStrategy{})

	lb.AddNode(Node{ID: "A", Address: "10.0.0.1:8080", ActiveConns: 10})
	lb.AddNode(Node{ID: "B", Address: "10.0.0.2:8080", ActiveConns: 5})
	lb.AddNode(Node{ID: "C", Address: "10.0.0.3:8080", ActiveConns: 20})

	for i := 0; i < 10; i++ {
		n := lb.GetNext()
		if n == nil {
			t.Fatalf("se esperaba un nodo, pero GetNext() devolvió nil en la iteración %d", i)
		}
		if n.ID != "B" {
			t.Fatalf("least connections eligió mal: se esperaba %q, obtenido %q", "B", n.ID)
		}
	}
}

func TestHotStrategySwap(t *testing.T) {
	lb := NewLoadBalancer(&RoundRobinStrategy{})

	// Hacemos que el mínimo sea obvio para LeastConn: B tiene 1 conexión
	lb.AddNode(Node{ID: "A", Address: "10.0.0.1:8080", ActiveConns: 10})
	lb.AddNode(Node{ID: "B", Address: "10.0.0.2:8080", ActiveConns: 1})
	lb.AddNode(Node{ID: "C", Address: "10.0.0.3:8080", ActiveConns: 20})

	const total = 100
	got := make([]string, 0, total)

	for i := 0; i < total; i++ {
		if i == total/2 {
			lb.SetStrategy(&LeastConnStrategy{})
		}

		n := lb.GetNext()
		if n == nil {
			t.Fatalf("se esperaba un nodo, pero GetNext() devolvió nil en la petición %d", i)
		}
		got = append(got, n.ID)
	}

	// La primera mitad debe ser Round Robin exacto con A,B,C repetidos
	wantFirstHalf := make([]string, 0, total/2)
	rrCycle := []string{"A", "B", "C"}
	for i := 0; i < total/2; i++ {
		wantFirstHalf = append(wantFirstHalf, rrCycle[i%len(rrCycle)])
	}

	for i := 0; i < total/2; i++ {
		if got[i] != wantFirstHalf[i] {
			t.Fatalf("antes del cambio de estrategia (petición %d) se esperaba %q, obtenido %q", i, wantFirstHalf[i], got[i])
		}
	}

	// 2) Segunda mitad debe ser LeastConn siempre B
	for i := total / 2; i < total; i++ {
		if got[i] != "B" {
			t.Fatalf("después del cambio de estrategia (petición %d) se esperaba %q, obtenido %q", i, "B", got[i])
		}
	}
}
