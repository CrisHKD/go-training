package dispatcher

import (
	"testing"
	"time"
)

// Genera N cantidad de tareas retornandolas en un slice tasks.
func generateTasks(n int, cost time.Duration) []Task {
	tasks := make([]Task, 0, n)
	for i := 0; i < n; i++ {
		tasks = append(tasks, Task{ID: i + 1, Cost: cost})
	}
	return tasks
}

func TestRunPool_ReceivesExactlyNResults(t *testing.T) {
	t.Parallel()

	const n = 100
	const workers = 3

	// costo pequeño para simular trabajo
	tasks := generateTasks(n, 2*time.Millisecond)

	results := RunPool(tasks, workers)
	if got := len(results); got != n {
		t.Fatalf("se esperaban %d resultados, pero se obtuvieron %d", n, got)
	}

	// Validación IDs únicos
	seen := make(map[int]struct{}, n)
	for _, r := range results {
		if r.Err != nil {
			t.Fatalf("error inesperado en la tarea %d: %v", r.TaskID, r.Err)
		}
		seen[r.TaskID] = struct{}{}
	}
	if len(seen) != n {
		t.Fatalf("se esperaban %d IDs de tareas únicos, pero se obtuvieron %d", n, len(seen))
	}
}

func BenchmarkSequential_100Tasks(b *testing.B) {
	tasks := generateTasks(100, 2*time.Millisecond)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RunSequential(tasks)
	}
}

func BenchmarkPool_3Workers_100Tasks(b *testing.B) {
	tasks := generateTasks(100, 2*time.Millisecond)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RunPool(tasks, 3)
	}
}
