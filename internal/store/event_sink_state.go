package store

type EventSinkResolver interface {
	Resolve(string) string
}

type event_sinkResolver struct{ prefix string }

func (r *event_sinkResolver) Resolve(key string) string {
	return r.prefix + key
}

func NewEventSinkResolver(enabled bool) EventSinkResolver {
	var resolver *event_sinkResolver
	if enabled {
		resolver = &event_sinkResolver{prefix: "ready:"}
	}
	return resolver
}
