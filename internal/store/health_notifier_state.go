package store

type HealthNotifierResolver interface {
	Resolve(string) string
}

type health_notifierResolver struct{ prefix string }

func (r *health_notifierResolver) Resolve(key string) string {
	return r.prefix + key
}

func NewHealthNotifierResolver(enabled bool) HealthNotifierResolver {
	var resolver *health_notifierResolver
	if enabled {
		resolver = &health_notifierResolver{prefix: "ready:"}
	}
	return resolver
}
