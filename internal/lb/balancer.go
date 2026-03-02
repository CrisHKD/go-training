package lb

type LoadBalancer struct {
	nodes    []Node
	strategy LoadBalancerStrategy
}

func NewLoadBalancer(strategy LoadBalancerStrategy) *LoadBalancer {
	return &LoadBalancer{
		nodes:    make([]Node, 0),
		strategy: strategy,
	}
}

func (lb *LoadBalancer) AddNode(n Node) {
	lb.nodes = append(lb.nodes, n)
}

func (lb *LoadBalancer) SetStrategy(s LoadBalancerStrategy) {
	lb.strategy = s
}

func (lb *LoadBalancer) GetNext() *Node {
	if len(lb.nodes) == 0 {
		return nil
	}
	if lb.strategy == nil {
		return nil
	}
	return lb.strategy.SelectNode(lb.nodes)
}
