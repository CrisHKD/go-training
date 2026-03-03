package provisioner

// Instance representa una máquina virtual creada en cualquier proveedor
// es una estructura común para que el sistema sea agnóstico a la nube
type Instance struct {
	ID     string // Identificador del servidor
	IP     string // Dirección IP asignada
	Status string // Estado actual del servidor
}

// CloudProvider define el contrato que debe cumplir cualquier proveedor de nube.
type CloudProvider interface {
	// CreateServer crea un servidor con el nombre dado.
	// Devuelve la instancia creada o un error si ocurre algún problema.
	CreateServer(name string) (*Instance, error)
}