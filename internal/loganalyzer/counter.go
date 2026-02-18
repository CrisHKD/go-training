package loganalyzer

import "strings"

func CountErrors(logs []string) map[string]int {

	counts := make(map[string]int)
	replacer := strings.NewReplacer("[", "", "]", "", ":", "")

	for _, log := range logs {
		clean := strings.TrimSpace(replacer.Replace(log))

		if clean == "" {
			continue
		}
		counts[clean]++
	}
	return counts
}

func CountErrorsPrealloc(logs []string, capacity int) map[string]int {
	replacer := strings.NewReplacer("[", "", "]", "", ":", "")

	if capacity < 0 {
		capacity = 0
	}
	counts := make(map[string]int, capacity)

	for _, raw := range logs {
		clean := strings.TrimSpace(replacer.Replace(raw))
		if clean == "" {
			continue
		}
		counts[clean]++
	}

	return counts
}
