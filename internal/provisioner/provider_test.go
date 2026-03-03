package provisioner

import "testing"

// Verifica que el factory cree correctamente un provider valido
func TestProviderCreation(t *testing.T) {
	p, err := GetProvider("aws")
	if err != nil {
		t.Fatalf("no se esperaba error, ocurrió: %v", err)
	}

	if p == nil {
		t.Fatalf("se esperaba un provider válido, pero fue nil")
	}

	// Probamos que realmente pueda crear un servidor
	inst, err := p.CreateServer("test-server")
	if err != nil {
		t.Fatalf("error creando servidor: %v", err)
	}

	if inst.IP == "" {
		t.Fatalf("la instancia no tiene IP asignada")
	}
}

// Verifica que el factory maneje correctamente un proveedor inexistente
func TestUnknownProvider(t *testing.T) {
	p, err := GetProvider("gcp")

	if err == nil {
		t.Fatalf("se esperaba error para proveedor desconocido")
	}

	if p != nil {
		t.Fatalf("se esperaba provider nil cuando hay error")
	}
}

// Prueba el flujo completo creando varios servidores
func TestProvisioningWorkflow(t *testing.T) {
	p, err := GetProvider("do")
	if err != nil {
		t.Fatalf("error obteniendo provider: %v", err)
	}

	names := []string{"api-1", "db-1", "worker-1"}

	for _, name := range names {
		inst, err := p.CreateServer(name)
		if err != nil {
			t.Fatalf("error creando servidor %s: %v", name, err)
		}

		if inst.IP == "" {
			t.Fatalf("servidor %s no tiene IP asignada", name)
		}

		if inst.Status != "running" {
			t.Fatalf("servidor %s no está en estado running", name)
		}
	}
}
