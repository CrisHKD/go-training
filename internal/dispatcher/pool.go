package dispatcher

import (
	"sync"
	"time"
)

// Estructura de un trabajo a ejecutar
type Task struct {
	ID   int
	Cost time.Duration
}

// Estructura de del resultado de procesar una tarea
type Result struct {
	TaskID int
	Err    error
}

// Worker realiza tareas desde jobs y guarda los resultados en results, termina cuando jobs se cierra
func Worker(jobs <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range jobs {
		if task.Cost > 0 {
			time.Sleep(task.Cost)
		}

		results <- Result{TaskID: task.ID, Err: nil}
	}
}

// RunPool ejecuta un pool de workers para procesar tasks
// Devuelve todos los resultados
func RunPool(tasks []Task, workerCount int) []Result {
	if workerCount <= 0 {
		workerCount = 1
	}

	jobs := make(chan Task)
	results := make(chan Result)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	// Inicia los workers
	for i := 0; i < workerCount; i++ {
		go Worker(jobs, results, &wg)
	}

	// Cierra results cuando todos los workers terminen
	go func() {
		wg.Wait()
		close(results)
	}()

	// Envia tareas y termina jobs
	go func() {
		for _, t := range tasks {
			jobs <- t
		}
		close(jobs)
	}()

	// Recolecta resultados hasta que results se cierre
	out := make([]Result, 0, len(tasks))

	for r := range results {
		out = append(out, r)
	}

	return out
}

// RunSequential procesa tareas de forma secuencial para benchmark
func RunSequential(tasks []Task) []Result {
	out := make([]Result, 0, len(tasks))

	for _, t := range tasks {
		if t.Cost > 0 {
			time.Sleep(t.Cost)
		}
		out = append(out, Result{TaskID: t.ID, Err: nil})
	}

	return out
}
