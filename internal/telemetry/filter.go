package telemetry

type Metric struct{
	ID string
	Value float64
}

func FilterHightUsage (data []Metric, threshold float64) []Metric{
	count := 0

	for _,m := range data {
		if m.Value >= threshold {
			count ++
		}
	}

	out := make([]Metric,0,count)

	for _,m := range data {
		if m.Value >= threshold {
			out = append(out, m)
		}
	}	
	return out
}