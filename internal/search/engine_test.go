package search

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

// GenerateFiles crea archivos n lineas.
// Inserta el patrón 2 veces por archivo en posiciones determinadas.
// Devuelve (files, expectedMatches).
func GenerateFiles(nFiles, nLines int, pattern string) (map[string][]string, int) {
	files := make(map[string][]string, nFiles)

	// Seed fijo para que tests sean reproducibles.
	rnd := rand.New(rand.NewSource(42))

	expexted := 0
	for f := 0; f < nFiles; f++ {
		name := fmt.Sprintf("file_%02d.log", f)
		lines := make([]string, nLines)

		for i := 0; i < nLines; i++ {
			lines[i] = randomLine(rnd)
		}

		// Insertamos el patrón en posiciones conocidas (2 por archivo).
		pos1 := (f*3 +1) % nLines
		pos2 := (f*5 +7) % nLines

		lines[pos1] = lines[pos1] + " " + pattern
		lines[pos2] = pattern + " " + lines[pos2]
		expexted += 2

		files[name] = lines
	}

	return files, expexted
}

func randomLine(rnd *rand.Rand) string{
	words := []string{"INFO", "WARN", "DEBUG", "servicio", "db", "cache", "timeout", "usuario", "reques", "ok"}
	n := 5 + rnd.Intn(8)

	var b strings.Builder
	for i := 0; i < n; i++{
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(words[rnd.Intn(len(words))])
	}
	return b.String()
}

func TestSearchAll_Precision(t *testing.T) {
	pattern := "CRITICAL_ERROR"
	files, expected := GenerateFiles(10, 50, pattern)

	results := SearchAll(pattern, files)

	if len(results) != expected {
		t.Fatalf("se esperaban %d resultados, pero se obtuvieron %d", expected, len(results))
	}

	// Validación extra: todas las coincidencias realmente contienen el patrón.
	for _, r := range results {
		if !strings.Contains(r.Match, pattern) {
			t.Fatalf("el resultado no contiene el patron esperado: %+v", r)
		}
	}
}

func TestSearchAll_OrderIndependent(t *testing.T) {
	// Este test asegura que NO dependemos del orden (concurrencia).
	pattern := "CRITICAL_ERROR"
	files, expected := GenerateFiles(10, 50, pattern)

	results := SearchAll(pattern, files)

	if len(results) != expected {
		t.Fatalf("se esperaban %d resultados, pero se obtuvieron %d", expected, len(results))
	}

	// Creamos un set con "file:line" encontrados.
	found := make(map[string]struct{}, len(results))
	for _, r := range results {
		key := fmt.Sprintf("%s:%d", r.File, r.Line)
		found[key] = struct{}{}
	}

	// Recalculamos exactamente las posiciones esperadas y verificamos que estén.
	for f := 0; f < 10; f++ {
		name := fmt.Sprintf("file_%02d.log", f)
		pos1 := (f*3 + 1) % 50
		pos2 := (f*5 + 7) % 50

		k1 := fmt.Sprintf("%s:%d", name, pos1)
		k2 := fmt.Sprintf("%s:%d", name, pos2)

		if _, ok := found[k1]; !ok {
			t.Fatalf("falta la coincidencia esperada en: %s", k1)
		}
		if _, ok := found[k2]; !ok {
			t.Fatalf("falta la coincidencia esperada en: %s", k2)
		}
	}
}