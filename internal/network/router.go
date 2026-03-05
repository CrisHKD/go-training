package network

type SmartRouter struct {
	routes map[string]string
}

// Constructor para inicializar el mapa de rutas
func NewSmartRouter() *SmartRouter {
	return &SmartRouter{
		routes: make(map[string]string),
	}
}

// Forward implementa la interfaz PacketForwarder, simula el envio de un paquete
func (r *SmartRouter) Forward(packet []byte) error {
	// En un router real aquí se decidiría la ruta del paquete
	return nil
}

// AddRoute implementa la interfaz RouteManager, agrega una nueva ruta a la tabla
func (r *SmartRouter) AddRoute(dest string, nextHop string) {
	r.routes[dest] = nextHop
}
