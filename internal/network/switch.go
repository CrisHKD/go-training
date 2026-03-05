package network

type LegacySwitch struct{}

// Forward implementa la interfaz PacketForwarder simula el reenvio de un paquete
func (s *LegacySwitch) Forward(packet []byte) error {
	// En un switch real aqui se enviaria el paquete al puerto correcto
	return nil
}
