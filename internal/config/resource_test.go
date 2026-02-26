package config

import (
	"bytes"
	"math/rand"
	"reflect"
	"testing"
)

// Generador de configuraciones aleatorias.
// Algunos campos quedan en zero value (0 o "").
func generateRandomConfig(r *rand.Rand) ResourceConfig {
	c := ResourceConfig{}

	// 50% de probabilidad de asignar CPU
	if r.Intn(2) == 1 {
		c.CPU = r.Intn(16) + 1
	}

	// 50% de probabilidad de asignar Memory
	if r.Intn(2) == 1 {
		c.Memory = (r.Intn(8) + 1) * 512
	}

	// 50% de probabilidad de asignar Burst
	if r.Intn(2) == 1 {
		modes := []string{"soft", "hard"}
		c.Burst = modes[r.Intn(len(modes))]
	}

	return c
}

// Test que verifica que omitempty funcione correctamente.
func TestExportConfig_OmitEmpty(t *testing.T) {
	c := ResourceConfig{
		CPU:    0,
		Memory: 2048,
		Burst:  "",
	}

	data, err := ExportConfig(c)
	if err != nil {
		t.Fatalf("error inesperado al exportar: %v", err)
	}

	// CPU no debe aparecer
	if bytes.Contains(data, []byte("cpu_limit")) {
		t.Fatalf("cpu_limit no debería aparecer en el JSON: %s", string(data))
	}

	// Burst no debe aparecer
	if bytes.Contains(data, []byte("burst")) {
		t.Fatalf("burst no debería aparecer en el JSON: %s", string(data))
	}

	// Memory sí debe aparecer
	if !bytes.Contains(data, []byte("memory_limit")) {
		t.Fatalf("memory_limit debería aparecer en el JSON: %s", string(data))
	}
}

// Test de round-trip: Struct -> JSON -> Struct
func TestRoundTrip_Integrity(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	for i := 0; i < 100; i++ {
		original := generateRandomConfig(r)

		data, err := ExportConfig(original)
		if err != nil {
			t.Fatalf("error al exportar: %v", err)
		}

		parsed, err := ParseConfig(data)
		if err != nil {
			t.Fatalf("error al parsear: %v", err)
		}

		if !reflect.DeepEqual(original, parsed) {
			t.Fatalf("round-trip falló.\nOriginal: %+v\nParsed: %+v\nJSON: %s",
				original, parsed, string(data))
		}
	}
}
