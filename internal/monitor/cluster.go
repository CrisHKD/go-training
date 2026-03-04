package monitor

// ClusterMonitor se encarga de registrar observers y notificar eventos
type ClusterMonitor struct {
	observers []HealthObserver
}

// NewClusterMonitor crea un nuevo monitor vacío
func NewClusterMonitor() *ClusterMonitor {
	return &ClusterMonitor{
		observers: []HealthObserver{},
	}
}

// Register agrega un nuevo observer a la lista
func (m *ClusterMonitor) Register(o HealthObserver) {
	m.observers = append(m.observers, o)
}

// Notify envia el evento a todos los observers registrados
func (m *ClusterMonitor) Notify(e HealthEvent) {
	for _, observer := range m.observers {
		observer.OnHealthChange(e)
	}
}
