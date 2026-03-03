package provisioner

import (
	"fmt"
	"math/rand"
	"time"
)

// AWSProvider simula un proveedor tipo AWS EC2
type AWSProvider struct {
	// Aca podemos guardar config
}

// CreateServer simula la creación de una instancia EC2
func (p *AWSProvider) CreateServer(name string) (*Instance, error) {
	// Simular que AWS tarda en crear recursos
	time.Sleep(300 * time.Millisecond)

	// Generar valores simulados
	id := fmt.Sprintf("i-%06d", rand.Intn(1_000_000))
	ip := fmt.Sprintf("10.0.0.%d", rand.Intn(200)+2)

	inst := &Instance{
		ID:     id,
		IP:     ip,
		Status: "running",
	}

	return inst, nil
}