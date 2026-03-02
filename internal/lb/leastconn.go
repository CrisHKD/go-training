package lb

type LeastConnStrategy struct{}

func (l *LeastConnStrategy) SelectNode(nodes []Node) *Node {
	if len(nodes) == 0 {
		return nil
	}

	// Asumimos que el primero es el minimo
	minIndex := 0

	// Recorremos los nodos para encontrar el minimo
	for i := 1; i < len(nodes); i++ {
		if nodes[i].ActiveConns < nodes[minIndex].ActiveConns {
			minIndex = i
		}
	}

	return &nodes[minIndex]
}
