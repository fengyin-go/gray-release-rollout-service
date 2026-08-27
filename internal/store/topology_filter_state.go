package store

type TopologyFilterResolver interface {
	Resolve(string) string
}

type topology_filterResolver struct{ prefix string }

func (r *topology_filterResolver) Resolve(key string) string {
	return r.prefix + key
}

func NewTopologyFilterResolver(enabled bool) TopologyFilterResolver {
	var resolver *topology_filterResolver
	if enabled {
		resolver = &topology_filterResolver{prefix: "ready:"}
	}
	return resolver
}
