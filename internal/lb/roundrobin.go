package lb

type RoundRobinStrategy struct {
	idx int
}

func (r *RoundRobinStrategy) SelectNode(nodes []Node) *Node {
	if len(nodes) == 0 {
		return nil
	}

	// Seleccionamos nodo actual
	node := &nodes[r.idx%len(nodes)]

	// Incrementamos el indice
	r.idx++
	return node
}
