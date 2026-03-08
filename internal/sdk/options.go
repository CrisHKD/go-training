package sdk

import "time"

// Option es una función que modifica la configuración del cliente
type Option func(*HttpClient)

// Cambia el timeout del cliente
func WithTimeout(d time.Duration) Option {
	return func(c *HttpClient) {
		c.timeout = d
	}
}

// Cambia la cantidad de reintentos
func WithRetries(n int) Option {
	return func(c *HttpClient) {
		c.retries = n
	}
}

// Activa o desactiva el modo debug
func WithDebug(debug bool) Option {
	return func(c *HttpClient) {
		c.debug = debug
	}
}