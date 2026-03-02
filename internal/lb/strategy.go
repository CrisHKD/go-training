package lb

type Node struct {
	ID          string
	Address     string
	ActiveConns int
}

type LoadBalancerStrategy interface {
	SelectNode(nodes []Node) *Node
}
