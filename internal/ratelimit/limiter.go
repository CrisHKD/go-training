package ratelimit

import "time"

// Limiter controla cada cuanto tiempo se puede procesar una peticion.
// Usamos un time.Ticker para generar permisos en intervalos regulares.
type Limiter struct {
	ticker *time.Ticker
}

// NewLimiter crea un nuevo Limiter que recibe un intervalo que indica cada cuánto tiempo se permite una petición.
func NewLimiter(interval time.Duration) *Limiter {
	return &Limiter{
		ticker: time.NewTicker(interval),
	}
}

// ProcessRequest bloquea la ejecución hasta que el ticker emita la siguiente señal.
func (l *Limiter) ProcessRequest() {
	<-l.ticker.C
}

// Stop detiene el ticker para evitar que siga corriendo en segundo plano.
func (l *Limiter) Stop() {
	l.ticker.Stop()
}
