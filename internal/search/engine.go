package search

import (
	"strings"
	"sync"
)

// Result representa una coinsidencia encontrada en un archivo
type Result struct{
	File string
	Line int
	Match string
}

// SearchFile procesa un archivo y envía coincidencias al canal out.
func searchFile(fileName string, lines []string, pattern string, out chan <- Result, wg *sync.WaitGroup){
	defer wg.Done()

	for i, line := range lines{
		if strings.Contains(line, pattern){
			out <- Result{
				File: fileName,
				Line: i,
				Match: line,
			}
		}
	}
}

// SearchAll implementa Fan-out/Fan-in:
// Fan-out: lanza una goroutine por archivo
// Fan-in: recolecta todos los resultados desde un canal central
func SearchAll(pattern string, files map[string][]string) []Result{
	// Canal central de resultados
	out := make(chan Result, 64)

	var wg sync.WaitGroup

	//Fan-out lanza una goroutine por archivo
	for name, lines := range files {
		wg.Add(1)
		go searchFile(name, lines, pattern, out, &wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	// Fan-in recolecta resultados en un slice final
	results := make([]Result, 0)
	for r := range out {
		results = append(results, r)
	}

	return results
}