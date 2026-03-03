package provisioner

import (
	"fmt"
	"math/rand"
	"time"
)

// DigitalOceanProvider simula un proveedor tipo DigitalOcean
type DigitalOceanProvider struct {
	// Aca podemos guardar config
}

// CreateServer simula la creación de un Droplet en DigitalOcean.
func (p *DigitalOceanProvider) CreateServer(name string) (*Instance, error) {
	// Simular que tarda en crear el servidor
	time.Sleep(200 * time.Millisecond)

	// Generar valores simulados
	id := fmt.Sprintf("do-%06d", rand.Intn(1_000_000))
	ip := fmt.Sprintf("192.168.1.%d", rand.Intn(200)+2)

	inst := &Instance{
		ID:     id,
		IP:     ip,
		Status: "running",
	}

	return inst, nil
}