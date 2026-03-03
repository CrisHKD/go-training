package provisioner

import (
	"fmt"
)

// GetProvider devuelve el proveedor correcto según el tipo solicitado
// Aplica el patrón Factory para desacoplar la creación de proveedores
func GetProvider(providerType string) (CloudProvider, error) {
	switch providerType {
	case "aws":
		return &AWSProvider{}, nil
	case "do":
		return &DigitalOceanProvider{}, nil
	default:
		return nil, fmt.Errorf("proveedor desconocido: %s", providerType)
	}
}
