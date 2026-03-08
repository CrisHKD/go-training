package sdk

import "time"

// HttpClient representa al cliente de la API
type HttpClient struct {
	addr    string
	timeout time.Duration
	retries int
	debug   bool
}

func NewClient(addr string, opts ...Option) *HttpClient {

	// Configuracion por defecto
	client := &HttpClient{
		addr:    addr,
		timeout: 30 * time.Second,
		retries: 3,
		debug:   false,
	}

	// Aplicar las opciones recibidas
	for _, opt := range opts {
		opt(client)
	}

	return client
}
