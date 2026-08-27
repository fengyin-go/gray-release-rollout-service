package store

type BalancerResolverResolver interface {
	Resolve(string) string
}

type balancer_resolverResolver struct{ prefix string }

func (r *balancer_resolverResolver) Resolve(key string) string {
	return r.prefix + key
}

func NewBalancerResolverResolver(enabled bool) BalancerResolverResolver {
	var resolver *balancer_resolverResolver
	if enabled {
		resolver = &balancer_resolverResolver{prefix: "ready:"}
	}
	return resolver
}
