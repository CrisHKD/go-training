package telemetry

type Metric struct{
	ID string
	Value float64
}

// FilterHightUsage retorna un nuevo slice son los datos filtrados
func FilterHightUsage (data []Metric, threshold float64) []Metric{
	count := 0

	// Contamos para asignar una capasidad exacta
	for _,m := range data {
		if m.Value >= threshold {
			count ++
		}
	}

	// Creamos el slice nuevo con la capacidad exacta
	out := make([]Metric,0,count)

	// Agregamos los datos filtrados al slice 	
	for _,m := range data {
		if m.Value >= threshold {
			out = append(out, m)
		}
	}	
	return out
}