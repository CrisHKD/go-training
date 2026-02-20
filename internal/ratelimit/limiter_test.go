package ratelimit

import (
	"sync"
	"testing"
	"time"
)

func TestRateLimiter_FlujoConstante(t *testing.T) {
	lim := NewLimiter(100 * time.Millisecond)
	defer lim.Stop()

	// Simulamos 5 peticiones una tras terminar otra
	start := time.Now()
	for i := 0; i < 5; i++ {
		lim.ProcessRequest()
	}
	elapsed := time.Since(start)

	// Con 100ms por request, 5 requests deben tardar al menos 1s.
	min := 500 * time.Millisecond
	if elapsed < min {
		t.Fatalf("El tiempo total fue %v, se esperaba al menos %v", elapsed, min)
	}
}

func TestRateLimiter_Rafaga(t *testing.T) {
	lim := NewLimiter(100 * time.Millisecond)
	defer lim.Stop()

	const n = 10

	var wg sync.WaitGroup
	wg.Add(n)

	start := time.Now()

	// Simular rafaga de 10 peticiones instantáneas.
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			lim.ProcessRequest()
		}()
	}

	// Esperar a que terminen todas.
	wg.Wait()
	elapsed := time.Since(start)

	// Con 100ms por request, 10 requests deben tardar al menos 1s.
	min := n * 100 * time.Millisecond
	if elapsed < min {
		t.Fatalf("El tiempo total fue %v, se esperaba al menos %v", elapsed, min)
	}
}
