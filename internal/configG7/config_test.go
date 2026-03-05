package configg7

import (
	"bytes"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

// resetSingleton reinicia el estado global para que cada test empiece limpio.
func resetSingleton() {
	instance = nil
	once = sync.Once{}
}

func TestSingletonIdentity(t *testing.T) {
	resetSingleton()

	c1 := GetConfig()
	c2 := GetConfig()
	c3 := GetConfig()
	c4 := GetConfig()
	c5 := GetConfig()
	c6 := GetConfig()
	c7 := GetConfig()
	c8 := GetConfig()
	c9 := GetConfig()
	c10 := GetConfig()

	// Comparación por puntero, todos deben ser la misma instancia
	all := []*Config{c1, c2, c3, c4, c5, c6, c7, c8, c9, c10}
	base := all[0]
	for i, c := range all {
		if c != base {
			t.Fatalf("Se esperaba la misma instancia , pero la variable #%d apunta a otra dirección", i+1)
		}
	}
}

func TestConcurrentAccess(t *testing.T) {
	resetSingleton()

	// Capturamos stdout para contar cuantas veces se imprime "Cargando datos..."
	origStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("no se pudo crear pipe para capturar stdout: %v", err)
	}
	os.Stdout = w

	var wg sync.WaitGroup
	const n = 100
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			_ = GetConfig()
		}()
	}

	wg.Wait()

	// Cerramos writer y restauramos stdout
	_ = w.Close()
	os.Stdout = origStdout

	// Leemos todo lo que se imprimio
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	_ = r.Close()

	out := buf.String()
	count := strings.Count(out, "Cargando datos...")

	if count != 1 {
		t.Fatalf("esperaba que 'Cargando datos...' se imprima 1 vez, pero se imprimió %d veces.\nSalida:\n%s", count, out)
	}
}
