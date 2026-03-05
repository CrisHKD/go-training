package network

// PacketForwarder representa el data plane, solo necesita reenviar paquetes
type PacketForwarder interface {
	Forward(packet []byte) error
}

// RouteManager representa el control plane, solo necesita administrar rutas
type RouteManager interface {
	AddRoute(dest string, nextHop string)
}
