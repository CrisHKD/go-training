package config

import "encoding/json"

// ResourceConfig representa límites de recursos de infraestructura.
// En Go si un campo es int y su valor es 0, "omitempty" hará que ese campo NO aparezca en el JSON generado.
type ResourceConfig struct {
	CPU    int    `json:"cpu_limit,omitempty"`
	Memory int    `json:"memory_limit,omitempty"`
	Burst  string `json:"burst,omitempty"`
}

// ParseConfig convierte JSON a struct.
func ParseConfig(data []byte) (ResourceConfig, error) {
	var c ResourceConfig
	err := json.Unmarshal(data, &c)
	if err != nil {
		return ResourceConfig{}, err
	}
	return c, nil
}

// ExportConfig convierte struct a JSON.
func ExportConfig(c ResourceConfig) ([]byte, error) {
	return json.Marshal(c)
}
