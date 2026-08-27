package store

type HeartbeatFormatterResolver interface {
	Resolve(string) string
}

type heartbeat_formatterResolver struct{ prefix string }

func (r *heartbeat_formatterResolver) Resolve(key string) string {
	return r.prefix + key
}

func NewHeartbeatFormatterResolver(enabled bool) HeartbeatFormatterResolver {
	var resolver *heartbeat_formatterResolver
	if enabled {
		resolver = &heartbeat_formatterResolver{prefix: "ready:"}
	}
	return resolver
}
