package loganalyzer

import (
	"math/rand"
	"testing"
	"time"
)

func generateLogs(n int) []string {
	pool := []string{
		"ERROR:404",
		"WARN:500",
		"INFO:200",
		"[ERROR:404]",
		"[WARN:500]",
		"[INFO:200]",
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make([]string, 0, n)

	for i := 0; i < n; i++ {
		out = append(out, pool[r.Intn(len(pool))])
	}
	return out
}

func TestCountErrors_SumIsN(t *testing.T) {
	logs := generateLogs(1000)
	got := CountErrors(logs)

	sum := 0
	for _, v := range got {
		sum += v
	}

	if sum != 1000 {
		t.Fatalf("Suma esperada %d, suma obtenida %d", 1000, sum)
	}
}

func BenchmarkCountErrors_MapEmpty(b *testing.B) {
	logs := generateLogs(10_000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountErrors(logs)
	}
}

func BenchmarkCountErrors_MapPrealloc(b *testing.B) {
	logs := generateLogs(10_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountErrorsPrealloc(logs, 10_000)
	}
}
