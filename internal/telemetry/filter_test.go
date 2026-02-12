package telemetry

import (
	"math/rand"
	"strconv"
	"testing"
)

// 
func generateMetrics(n int) []Metric {
	data := make([]Metric, n)

	for i := 0; i < n; i++ {
		data[i] = Metric{
			ID:    "metric-" + strconv.Itoa(i),
			Value: rand.Float64() * 100.0,
		}
	}

	return data
}

func TestIndependencia(t *testing.T) {
	rand.Seed(34)

	original := generateMetrics(20)
	threshold := 50.0

	filtered := FilterHightUsage(original, threshold)

	if len(filtered) == 0 {
		original[0].Value = 51
		filtered = FilterHightUsage(original, threshold)
	}

	id := filtered[0].ID

	var sourceIdx = -1
	for i := range original {
		if original[i].ID == id {
			sourceIdx = i
			break
		}
	}

	if sourceIdx == -1 {
		t.Fatalf("no se encontro en el original el ID: %s", id)
	}

	before := original[sourceIdx].Value

	filtered[0].Value = before + 123.0

	if original[sourceIdx].Value != before {
		t.Fatalf("hubo comparticion: original cambio de %v a %v", before, original[sourceIdx].Value)
	}
}

func TestFiltrado_ConteoCorrecto(t *testing.T) {
	rand.Seed(99)

	original := generateMetrics(10)
	threshold := 70.0

	expected := 0
	for _, m := range original {
		if m.Value > threshold {
			expected++
		}
	}

	filtered := FilterHightUsage(original, threshold)

	if len(filtered) != expected {
		t.Fatalf("conteo incorrecto: esperado %d, obtenido %d", expected, len(filtered))
	}

	for _, m := range filtered {
		if m.Value <= threshold {
			t.Fatalf("se coló una métrica con Value=%v (threshold=%v)", m.Value, threshold)
		}
	}
}
